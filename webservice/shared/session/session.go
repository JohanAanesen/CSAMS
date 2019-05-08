package session

import (
	"fmt"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

// Private variables for this package
var (
	store *sessions.CookieStore
	name  string
)

// Session struct
type Session struct {
	Options   sessions.Options `json:"options"`
	Name      string           `json:"name"`
	SecretKey string           `json:"secretKey"`
}

// Configure Session
func Configure(s *Session) {
	store = sessions.NewCookieStore([]byte(s.SecretKey))
	store.Options = &s.Options
	name = s.Name
}

// Instance returns an instance of the Session
func Instance(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, name)
}

// Empty deletes all keys and values in the session
func Empty(sess *sessions.Session) {
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}

//IsTeacher returns if user is a teacher or not
func IsTeacher(r *http.Request) bool {
	//check if user is already logged in
	user := GetUserFromSession(r)

	//check that user is a teacher
	if !user.Teacher { //not a teacher or logged in
		return false
	}

	return IsLoggedIn(r)
}

//IsLoggedIn returns if user is authenticated or not
func IsLoggedIn(r *http.Request) bool {
	//get user from session
	user := GetUserFromSession(r)

	return user.Authenticated
}

//GetUserFromSession returns user object stored in session
func GetUserFromSession(r *http.Request) model.User {
	session, err := Instance(r) // get session

	if err != nil {
		log.Printf("ocould not get instance of session: %v", err)
		return model.User{Authenticated: false}
	}

	val := session.Values["user"]
	var user = model.User{}
	user, ok := val.(model.User)
	if !ok {
		return model.User{Authenticated: false}
	}
	return user
}

//SaveUserToSession saves user object to session store
func SaveUserToSession(user model.User, w http.ResponseWriter, r *http.Request) bool {
	session, err := Instance(r) //get session

	if err != nil {
		log.Printf("ocould not get instance of session: %v", err)
		return false
	}

	session.Values["user"] = user

	err = session.Save(r, w) //save session changes

	if err != nil {
		log.Printf("ocould save session: %v", err)
		//redirect somewhere
		return false
	}

	return true
}

func SetFlash(message string, w http.ResponseWriter, r *http.Request) error {
	session, err := Instance(r)
	if err != nil {
		return err
	}

	session.AddFlash(message, "message")
	return session.Save(r, w)
}

func GetFlash(w http.ResponseWriter, r *http.Request) string {
	session, err := Instance(r)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	fm := session.Flashes("message")
	if fm == nil {
		return ""
	}

	_ = session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}

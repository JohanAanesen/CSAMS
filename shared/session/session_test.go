package session_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/controller"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSveinNewSession(t *testing.T) {
	test := struct {
		expectedName      string
		expectedSecretKey string
		expectedOptions   *sessions.Options
	}{
		expectedName:      "gosess",
		expectedSecretKey: "52 67 166 253 96 202 151 106 65 44 177 84 130 1 208 172 233 228 151 112 132 236 225 112 168 222 202 121 102 43 41 151 54 129 105 1 233 5 77 68 207 10 251 15 252 134 240 64 171 237 177 154 209 203 62 3 116 138 74 175 97 177 16 156",
		expectedOptions: &sessions.Options{
			Path: "/",
			Domain: "",
			MaxAge: 28800,
			Secure: false,
			HttpOnly: true,
		},
	}

	cfg, _ := config.Load("../../config/config.json")
	sess := cfg.Session

	if sess.Name != test.expectedName {
		t.Logf("\nName:\n\t- expected: %v\n\t- got: %v\n", test.expectedName, sess.Name)
		t.Fail()
	}

	if sess.SecretKey != test.expectedSecretKey {
		t.Logf("\nSecretKey:\n\t- expected: %v\n\t- got: %v\n", test.expectedSecretKey, sess.SecretKey)
		t.Fail()
	}

	if sess.Options.Path != test.expectedOptions.Path {
		t.Logf("\nOptions.Path:\n\t- expected: %v\n\t- got: %v\n", test.expectedOptions.Path, sess.Options.Path)
		t.Fail()
	}

	if sess.Options.Domain != test.expectedOptions.Domain {
		t.Logf("\nOptions.Domain:\n\t- expected: %v\n\t- got: %v\n", test.expectedOptions.Domain, sess.Options.Domain)
		t.Fail()
	}

	if sess.Options.MaxAge != test.expectedOptions.MaxAge {
		t.Logf("\nOptions.MaxAge:\n\t- expected: %v\n\t- got: %v\n", test.expectedOptions.MaxAge, sess.Options.MaxAge)
		t.Fail()
	}

	if sess.Options.Secure != test.expectedOptions.Secure {
		t.Logf("\nOptions.Secure:\n\t- expected: %v\n\t- got: %v\n", test.expectedOptions.Secure, sess.Options.Secure)
		t.Fail()
	}

	if sess.Options.HttpOnly != test.expectedOptions.HttpOnly {
		t.Logf("\nOptions.HttpOnly:\n\t- expected: %v\n\t- got: %v\n", test.expectedOptions.HttpOnly, sess.Options.HttpOnly)
		t.Fail()
	}
}

func TestGetUserFromSession(t *testing.T) {
	if err := os.Chdir("../../"); err != nil { //go out of /handlers folder
		panic(err)
	}

	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	sess, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		ID:            id,
		Name:          name,
		EmailStudent:  email,
	}

	//save user to session values
	sess.Values["user"] = user
	//save session changes
	err = sess.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(controller.IndexGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	user2 := session.GetUserFromSession(req)

	if user2.ID != id {
		t.Errorf("Returned wrong user information from session, expected %v, got %v", id, user2.ID)
	}
	if user2.Name != name {
		t.Errorf("Returned wrong user information from session, expected %v, got %v", name, user2.Name)
	}
	if user2.EmailStudent != email {
		t.Errorf("Returned wrong user information from session, expected %v, got %v", email, user2.EmailStudent)
	}
}

func TestIsLoggedIn(t *testing.T) {

	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	sess, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		ID:            id,
		Name:          name,
		EmailStudent:  email,
	}

	//save user to session values
	sess.Values["user"] = user
	//save session changes
	err = sess.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(controller.IndexGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	loggedIn := session.IsLoggedIn(req)

	if !loggedIn {
		t.Errorf("Not logged in expected true, got false")
	}
}

func TestIsTeacher(t *testing.T) {
	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	sess, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       true,
		ID:            id,
		Name:          name,
		EmailStudent:  email,
	}

	//save user to session values
	sess.Values["user"] = user
	//save session changes
	err = sess.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(controller.IndexGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	isTeacher := session.IsTeacher(req)

	if !isTeacher {
		t.Errorf("Not logged in expected true, got false")
	}
}

func TestSaveUserToSession(t *testing.T) {
	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		ID:            id,
		Name:          name,
		EmailStudent:  email,
	}

	http.HandlerFunc(controller.IndexGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	session.SaveUserToSession(user, resp, req)

	user2 := session.GetUserFromSession(req)

	if user2.ID != id {
		t.Errorf("Returned wrong user information from session, expected %v, got %v", id, user2.ID)
	}
	if user2.Name != name {
		t.Errorf("Returned wrong user information from session, expected %v, got %v", name, user2.Name)
	}
	if user2.EmailStudent != email {
		t.Errorf("Returned wrong user information from session, expected %v, got %v", email, user2.EmailStudent)
	}
}

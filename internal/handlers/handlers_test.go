package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	db.InitDB(os.Getenv("SQLDB"))

	if err := os.Chdir("../../"); err != nil { //go out of /handlers folder
		panic(err)
	}
}

func TestLoginHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(LoginHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestLoggingIn(t *testing.T) {
	/*
	form := url.Values{}
	form.Add("email", "hei@gmail.com")
	form.Add("password", "hei")
	req := httptest.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	req.Form = form

	resp := httptest.NewRecorder()

	http.HandlerFunc(LoginRequest).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
	*/
}

func TestMainHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestCourseHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/course?id=asdfcvgbhnjk", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(CourseHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestCourseListHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/course/list?id=adsikjuh", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(CourseListHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestUserHandler(t *testing.T) {

	// TODO : fix this
	/*
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(UserHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
	*/
}

// First test that the user gets redirected to /login if he's not logged in
func TestCheckUserStatusNotLoggedIn(t *testing.T) {

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(UserHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusFound {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusFound, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestCheckUserStatusLoggedIn(t *testing.T) {
	// TODO : Fix code
}

func TestUserUpdateRequest(t *testing.T) {
	// TODO : Fix code

	/*
	// Change name and email
	form := url.Values{}
	form.Add("name", "Name N. Nameson")
	form.Add("secondaryEmail", "myNew@emailIsCool.com")
	req := httptest.NewRequest("POST", "/user", strings.NewReader(form.Encode()))
	req.Form = form

	resp := httptest.NewRecorder()

	http.HandlerFunc(UserUpdateRequest).ServeHTTP(resp, req)

	status := resp.Code

	// Status should be 200/OK if it went by okey
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
	*/
}

func TestAdminHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	session, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user model.User
	user.Authenticated = true
	user.Teacher = true
	//save user to session values
	session.Values["user"] = user
	//save session changes
	err = session.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(AdminHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAdminCourseHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/course", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	session, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user model.User
	user.Authenticated = true
	user.Teacher = true
	//save user to session values
	session.Values["user"] = user
	//save session changes
	err = session.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(AdminCourseHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAdminCreateCourseHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/course/create", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	session, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user model.User
	user.Authenticated = true
	user.Teacher = true
	//save user to session values
	session.Values["user"] = user
	//save session changes
	err = session.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(AdminCreateCourseHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAdminUpdateCourseHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/course/update", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	session, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user model.User
	user.Authenticated = true
	user.Teacher = true
	//save user to session values
	session.Values["user"] = user
	//save session changes
	err = session.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(AdminUpdateCourseHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAssignmentHandlerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment?id=ihadls&class=asdbjlid", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAssignmentAutoHandlerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment/auto?id=ihadls&class=asdbjlid", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentAutoHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAssignmentPeerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment/peer?id=ihadls&class=asdbjlid", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentPeerHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestErrorHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/course", nil) //class with no id should give 403
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentPeerHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusBadRequest, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestLogoutHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(LogoutHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusFound {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusFound, status)
	}
}

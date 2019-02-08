package handlers

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
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

	if status != http.StatusFound { //todo update this somehow, site isn't available for unauthorized users
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusFound, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestClassHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/class?id=asdfcvgbhnjk", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(ClassHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestClassListHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/class/list?id=adsikjuh", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(ClassListHandler).ServeHTTP(resp, req)

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
func TestCheckUserStatusNotLoggedIn(t *testing.T){

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
	req, err := http.NewRequest("GET", "/class", nil) //class with no id should give 403
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

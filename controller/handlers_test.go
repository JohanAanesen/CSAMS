package controller_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/controller"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func init() {
	if err := os.Chdir("../"); err != nil { //go out of /handlers folder
		panic(err)
	}

	var cfg = &config.Configuration{}
	cfg, _ = config.Load("config/config.json")

	// Configure Session
	session.Configure(cfg.Session)

	// Configure Database
	db.ConfigureDB(cfg.Database)

	db.OpenDB()
	defer db.CloseDB()
}

func TestHandlers(t *testing.T) {
	config.Initialize()

	tests := []struct {
		name    string
		method  string
		url     string
		body    io.Reader
		handler func(w http.ResponseWriter, r *http.Request)

		expectedCode int
	}{
		{
			name:         "index",
			method:       "GET",
			url:          "/",
			body:         nil,
			handler:      controller.IndexGET,
			expectedCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		r, _ := http.NewRequest(test.method, test.url, test.body)
		w := httptest.NewRecorder()

		test.handler(w, r)

		t.Run(test.name, func(t *testing.T) {
			if w.Body.String() == "" {
				t.Logf("error, response body was empty")
				t.Fail()
			}

			// Check status code
			if w.Code != test.expectedCode {
				t.Logf("expected: %v, got: %v\n", http.StatusOK, w.Code)
				t.Fail()
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	config.Initialize()

	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(controller.LoginGET).ServeHTTP(resp, req)

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
	config.Initialize()

	form := url.Values{}
	form.Add("email", "hei@gmail.com")
	form.Add("password", "hei")
	req := httptest.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	req.Form = form

	resp := httptest.NewRecorder()

	http.HandlerFunc(controller.LoginPOST).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusFound {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusFound, status)
	}

	body := resp.Body

	if body.Len() != 0 { //should be 0 because redirect if success
		t.Errorf("Response body error, expected 0, got %d", body.Len())
	}
}

func TestRegisterGET(t *testing.T) {
	config.Initialize()

	req, err := http.NewRequest("GET", "/register", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(controller.RegisterGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

}

func TestRegisterPOST(t *testing.T) {
	/*
		config.Initialize()

		form := url.Values{}
		form.Add("name", "Swag Meister")
		form.Add("email", "New@guy.com")
		form.Add("password", "hei")
		form.Add("passwordConfirm", "hei")
		req := httptest.NewRequest("POST", "/register", nil)
		req.Form = form

		resp := httptest.NewRecorder()

		http.HandlerFunc(controller.RegisterPOST).ServeHTTP(resp, req)

		status := resp.Code

		if status != http.StatusFound {
			t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusFound, status)
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

	http.HandlerFunc(controller.IndexGET).ServeHTTP(resp, req)

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

	req, err := http.NewRequest("GET", "/course?id=1", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(controller.CourseGET).ServeHTTP(resp, req)

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

	http.HandlerFunc(controller.CourseListGET).ServeHTTP(resp, req)

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

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       false,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
	}

	http.HandlerFunc(controller.UserGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

// First test that the user gets redirected to /login if he's not logged in
func TestCheckUserStatusNotLoggedIn(t *testing.T) {

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(controller.UserGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusUnauthorized, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestCheckUserStatusLoggedIn(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       false,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
	}

	http.HandlerFunc(controller.UserGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestUserUpdateRequest(t *testing.T) {
	// TODO : Fix code
	// Change name and email
	form := url.Values{}
	form.Add("usersName", "Ken Thompson") // One of the Golang creators
	form.Add("secondaryEmail", "mannen@harmannenfalt.no")
	req := httptest.NewRequest("POST", "/user/update", strings.NewReader(form.Encode()))
	req.Form = form

	resp := httptest.NewRecorder()

	//user object we want to fill with variables needed
	user := model.User{
		ID:            1,
		Name:          "Test User",
		EmailStudent:  "hei@gmail.com",
		EmailPrivate:  "test@yahoo.com",
		Authenticated: true,
		Teacher:       true,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
	}

	http.HandlerFunc(controller.UserUpdatePOST).ServeHTTP(resp, req)

	status := resp.Code

	// Status should be 200/OK if it went by okey
	if status != http.StatusFound {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusFound, status)
	}

	body := resp.Body

	if body.Len() != 0 { //should be 0 because of redirect if successfull
		t.Errorf("Response body error, expected 0, got %d", body.Len())
	}
}

func TestAdminHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       true,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
	}

	http.HandlerFunc(controller.AdminGET).ServeHTTP(resp, req)

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

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       true,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
	}

	http.HandlerFunc(controller.AdminCourseGET).ServeHTTP(resp, req)

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

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       true,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
	}

	http.HandlerFunc(controller.AdminCreateCourseGET).ServeHTTP(resp, req)

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

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       true,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
	}

	http.HandlerFunc(controller.AdminUpdateCourseGET).ServeHTTP(resp, req)

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

	http.HandlerFunc(controller.AssignmentGET).ServeHTTP(resp, req)

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

	http.HandlerFunc(controller.AssignmentAutoGET).ServeHTTP(resp, req)

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

	http.HandlerFunc(controller.AssignmentPeerGET).ServeHTTP(resp, req)

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

	http.HandlerFunc(controller.AssignmentPeerGET).ServeHTTP(resp, req)

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

	http.HandlerFunc(controller.LogoutGET).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusFound {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusFound, status)
	}
}

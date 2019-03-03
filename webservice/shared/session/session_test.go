package session_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/controller"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSession(t *testing.T) {
	config.Initialize("../../config/config.json")

	tests := []struct {
		name string
		f    func(t *testing.T)
	}{
		{
			name: "TestGetUserFromSession",
			f:    TestGetUserFromSession,
		},
		{
			name: "TestGetUserFromSession",
			f:    TestGetUserFromSession,
		},
		{
			name: "TestIsLoggedIn",
			f:    TestIsLoggedIn,
		},
		{
			name: "TestIsTeacher",
			f:    TestIsTeacher,
		},
		{
			name: "TestSaveUserToSession",
			f:    TestSaveUserToSession,
		},
		{
			name: "TestEmpty",
			f:    TestEmpty,
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.f)
	}
}

func TestEmpty(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatalf("new request: %s", err)
	}

	sess, err := session.Instance(r)
	if err != nil {
		log.Fatalf("session instance: %s", err)
	}

	sess.Values["foo"] = "bar"

	session.Empty(sess)

	if sess.Values["foo"] == "bar" {
		t.Fail()
	}
}

func TestGetUserFromSession(t *testing.T) {
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

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
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

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		ID:            id,
		Name:          name,
		EmailStudent:  email,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
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

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated: true,
		Teacher:       true,
		ID:            id,
		Name:          name,
		EmailStudent:  email,
	}

	//save user to session
	if !session.SaveUserToSession(user, resp, req) {
		t.Error("failed to save user to session")
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

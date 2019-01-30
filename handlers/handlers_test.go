package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	if err := os.Chdir("../"); err != nil { //go out of /handlers folder
		panic(err)
	}
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
}

func TestLoginHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestLogoutHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestClassHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/class", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestClassListHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/class/list", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestUserHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestAdminHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/admin", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestAssignmentHandlerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestAssignmentAutoHandlerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment/auto", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestAssignmentPeerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment/peer", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}
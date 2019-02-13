package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"net/http"
	"strings"
)

// ErrorHandler handles displaying errors for the clients
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	v := view.New(r)
	v.Name = "error"

	v.Vars["ErrorCode"] = status
	v.Vars["ErrorMessage"] = http.StatusText(status)

	v.Render(w)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	var admin string
	if strings.HasPrefix(r.RequestURI, "/admin") {
		admin = "admin/"
	}

	v := view.New(r)
	v.Name = admin + "error"

	v.Vars["ErrorCode"] = http.StatusNotFound
	v.Vars["ErrorMessage"] = http.StatusText(http.StatusNotFound)

	v.Render(w)
}
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusMethodNotAllowed)

	var admin string
	// Check if we it is a
	if strings.HasPrefix(r.RequestURI, "/admin") {
		admin = "admin/"
	}

	v := view.New(r)
	v.Name = admin + "error"

	v.Vars["ErrorCode"] = http.StatusMethodNotAllowed
	v.Vars["ErrorMessage"] = http.StatusText(http.StatusMethodNotAllowed)

	v.Render(w)
}

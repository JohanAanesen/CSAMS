package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"net/http"
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

// NotFoundHandler ... TODO (Svein) add comment here
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	v := view.New(r)
	v.Name = "error"

	v.Vars["ErrorCode"] = http.StatusNotFound
	v.Vars["ErrorMessage"] = http.StatusText(http.StatusNotFound)

	v.Render(w)
}

// MethodNotAllowedHandler ... TODO (Svein) add comment here
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusMethodNotAllowed)

	v := view.New(r)
	v.Name = "error"

	v.Vars["ErrorCode"] = http.StatusMethodNotAllowed
	v.Vars["ErrorMessage"] = http.StatusText(http.StatusMethodNotAllowed)

	v.Render(w)
}

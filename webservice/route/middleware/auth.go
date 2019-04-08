package middleware

import (
	"github.com/JohanAanesen/CSAMS/webservice/controller"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"net/http"
)

// TeacherAuth check on all it's request if the user is authorized as a teacher
func TeacherAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// User is not a teacher
		if !session.IsTeacher(r) {
			// Redirect to front-page
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}
		// Serve next respond and request
		next.ServeHTTP(w, r)
	})
}

// UserAuth check if user is authenticated
func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := session.GetUserFromSession(r)

		if r.RequestURI == "/" {
			switch r.Method {
			case "POST":
				controller.JoinCoursePOST(w, r)
				return
			}
		}
		// User is not authenticated
		if !currentUser.Authenticated {
			// Redirect to login
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		// Serve next respond and request
		next.ServeHTTP(w, r)
	})
}

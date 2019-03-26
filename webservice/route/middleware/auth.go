package middleware

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/controller"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"net/http"
)

// TeacherAuth check on all it's request if the user is authorized as a teacher
func TeacherAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsTeacher(r) { //not a teacher, error 401
			controller.ErrorHandler(w, r, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// UserAuth check if user is authenticated
func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := session.GetUserFromSession(r)

		if r.RequestURI == "/" {
			switch r.Method {
			case "GET":
				controller.IndexGET(w, r)
				return
			case "POST":
				controller.JoinCoursePOST(w, r)
				return
			}
		}

		if !currentUser.Authenticated {
			controller.ErrorHandler(w, r, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
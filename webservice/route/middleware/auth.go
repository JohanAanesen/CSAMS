package middleware

import (
	"net/http"
)

// TeacherAuth check on all it's request if the user is authorized as a teacher
func TeacherAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO (Svein): Fix this. Session broke down?
		/*if !session.IsTeacher(r) { //not a teacher, error 401
			controller.ErrorHandler(w, r, http.StatusUnauthorized)
			return
		}*/

		next.ServeHTTP(w, r)
	})
}

package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"net/http"
)

// IndexGET serves homepage to authenticated users, send anonymous to login
func IndexGET(w http.ResponseWriter, r *http.Request) {
	auth := session.GetUserFromSession(r).Authenticated

	if !auth {
		LoginGET(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// TODO (Svein): Add courses in v.Vars["Courses"]

	v := view.New(r)
	v.Name = "index"
	v.Vars["Auth"] = auth
	v.Render(w)
}

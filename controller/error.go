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

	/*
	temp, err := template.ParseFiles("template/layout.html", "template/navbar.html", "template/error.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle   string
		LoadFormCSS bool

		ErrorCode    int
		ErrorMessage string

		Menu model.Menu
	}{
		PageTitle:   "Error: " + string(status),
		LoadFormCSS: true,

		ErrorCode:    status,
		ErrorMessage: http.StatusText(status),

		Menu: util.LoadMenuConfig("configs/menu/site.json"),
	}); err != nil {
		log.Fatal(err)
	}
	*/
}

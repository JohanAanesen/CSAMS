package controller

import (
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"log"
	"net/http"
)

// AdminLogsGet serves the logs page
func AdminLogsGet(w http.ResponseWriter, r *http.Request) {

	// Services
	services := service.NewServices(db.GetDB())

	// Get logs
	logs, err := services.Logs.FetchAll()
	if err != nil {
		log.Println("services, logs, fetchall", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Set header content type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	// Set template file
	v.Name = "admin/logs/index"
	// Set variables
	v.Vars["Logs"] = logs
	// Render view
	v.Render(w)
}

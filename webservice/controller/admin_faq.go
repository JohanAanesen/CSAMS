package controller

import (
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"log"
	"net/http"
)

// AdminFaqGET handles GET-request at admin/faq/index
func AdminFaqGET(w http.ResponseWriter, r *http.Request) {

	// TODO (Svein): Move this to 'settings'
	// TODO (Svein): Allow blank FAQ

	// Services
	services := service.NewServices(db.GetDB())

	// Get current faq, date and questions
	content, err := services.FAQ.Fetch()
	if err != nil {
		log.Println("services, faq, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/faq/index"
	v.Vars["Updated"] = content

	v.Render(w)
}

// AdminFaqEditGET returns the edit view for the faq
func AdminFaqEditGET(w http.ResponseWriter, r *http.Request) {

	// TODO (Svein): Move this to 'settings'
	// TODO (Svein): Allow blank FAQ

	// Services
	services := service.NewServices(db.GetDB())

	// Get current faq, date and questions
	content, err := services.FAQ.Fetch()
	if err != nil {
		log.Println("services, faq, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/faq/edit"
	v.Vars["Content"] = content

	v.Render(w)
}

// AdminFaqUpdatePOST handles the edited markdown faq
func AdminFaqUpdatePOST(w http.ResponseWriter, r *http.Request) {
	// Check that the questions arrived
	updatedFAQ := r.FormValue("rawQuestions")
	if updatedFAQ == "" {
		log.Println("Form is empty!")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Get user for logging purposes
	currentUser := session.GetUserFromSession(r)

	// Get current unchanged faq
	content, err := services.FAQ.Fetch()
	if err != nil {
		log.Println("services, faq, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that it's changes to the new faq
	if content.Questions == updatedFAQ {
		log.Println("Old and new faq can not be equal!")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Update faq questions
	err = services.FAQ.Update(updatedFAQ)
	if err != nil {
		log.Println("services, faq, update", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Log update faq to db
	err = services.Logs.InsertUpdateFAQ(currentUser.ID, content.Questions, updatedFAQ)
	if err != nil {
		log.Println("log, update faq ", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//AdminFaqGET(w, r)
	http.Redirect(w, r, "/admin/faq", http.StatusFound)
}

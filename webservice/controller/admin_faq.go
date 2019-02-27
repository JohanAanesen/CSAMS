package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
)

// AdminFaqGET handles GET-request at admin/faq/index
func AdminFaqGET(w http.ResponseWriter, r *http.Request) {
	content := model.GetDateAndQuestionsFAQ()
	if content.Questions == "-1" {
		log.Println("Something went wrong with getting the faq (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Convert to html
	questions := github_flavored_markdown.Markdown([]byte(content.Questions))

	v := view.New(r)
	v.Name = "admin/faq/index"
	v.Vars["Updated"] = content.Date.Format("02. January 2006 - 15:04")
	v.Vars["Questions"] = template.HTML(questions)

	v.Render(w)
}

// AdminFaqEditGET returns the edit view for the faq
func AdminFaqEditGET(w http.ResponseWriter, r *http.Request) {
	//
	content := model.GetDateAndQuestionsFAQ()
	if content.Questions == "-1" {
		log.Println("Something went wrong with getting the faq (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/faq/edit"
	v.Vars["Updated"] = content.Date.Format("02. January 2006 - 15:04")
	v.Vars["RawContent"] = content.Questions

	v.Render(w)
}

// AdminFaqUpdatePOST handles the edited markdown faq
func AdminFaqUpdatePOST(w http.ResponseWriter, r *http.Request) {
	// Check that the questions arrived
	updatedFAQ := r.FormValue("rawQuestions")
	if updatedFAQ == "" {
		log.Println("Form is empty! (admin.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Check that it's possible to get the old faq from db
	content := model.GetDateAndQuestionsFAQ()
	if content.Questions == "-1" {
		log.Println("Something went wrong with getting the faq (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that it's changes to the new faq
	if content.Questions == updatedFAQ {
		log.Println("Old and new faq can not be equal! (admin.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Check that it went okay to add new faq to db
	if !model.UpdateFAQ(updatedFAQ) {
		log.Println("Something went wrong with updating the faq! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get user for logging purposes
	user := session.GetUserFromSession(r)

	// Collect the log data
	logData := model.Log{
		UserID:   user.ID,
		Activity: model.UpdateAdminFAQ,
		OldValue: content.Questions,
		NewValue: updatedFAQ,
	}

	// Log that a teacher has changed the faq
	if !model.LogToDB(logData) {
		log.Println("Something went wrong with logging the new faq! (admin.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	AdminFaqGET(w, r)
}

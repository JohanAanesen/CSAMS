package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"log"
	"net/http"
)

// AdminReviewGET handles GET-requests @ /admin/review
func AdminReviewGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	var reviewRepo = model.ReviewRepository{}

	submissions, err := reviewRepo.GetAll()
	if err != nil {
		log.Println(err)
		return
	}

	v := view.New(r)
	v.Name = "admin/submission/index"

	v.Vars["Submissions"] = submissions

	v.Render(w)
}

// AdminReviewCreateGET handles GET-requests @ /admin/review/create
func AdminReviewCreateGET(w http.ResponseWriter, r *http.Request) {

}

// AdminReviewCreatePOST handles POST-requests @ /admin/review/create
func AdminReviewCreatePOST(w http.ResponseWriter, r *http.Request) {

}

// AdminReviewUpdateGET handles GET-requests @ /admin/review/update/{id:[0-9]+}
func AdminReviewUpdateGET(w http.ResponseWriter, r *http.Request) {

}

// AdminReviewUpdatePOST handles POST-requests @ /admin/review/update
func AdminReviewUpdatePOST(w http.ResponseWriter, r *http.Request) {

}
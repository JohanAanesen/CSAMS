package plugin

import (
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"html/template"
)

// Review create template function maps related to reviews
func Review() template.FuncMap {
	f := make(template.FuncMap)

	services := service.NewServices(db.GetDB())

	f["HasBeenReviewed"] = func(targetID, reviewerID, assignmentID int) bool {
		check, err := services.ReviewAnswer.HasBeenReviewed(targetID, reviewerID, assignmentID)
		if err != nil {
			return false
		}

		return check
	}

	return f
}

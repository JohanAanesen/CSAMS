package controller

import (
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"net/http"
)

//PrivacyGET serves privay policy page
func PrivacyGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "privacy-policy"

	v.Render(w)
}

package controller

import (
	"github.com/kongebra/awesomeProject/shared/view"
	"net/http"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "index/index"
	v.Render(w)
}

package plugin

import (
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"html/template"
)

// GetUsername is a template function that converts user id to username
func GetUsername() template.FuncMap {
	f := make(template.FuncMap)

	f["GET_USERNAME"] = func(id int) string {
		user := model.GetUser(id)
		return user.Name
	}

	return f
}

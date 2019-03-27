package plugin

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
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

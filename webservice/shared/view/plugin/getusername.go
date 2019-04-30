package plugin

import (
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"html/template"
)

// GetUsername is a template function that converts user id to username
func GetUsername() template.FuncMap {
	//Services
	services := service.NewServices(db.GetDB())

	f := make(template.FuncMap)

	f["GET_USERNAME"] = func(id int) string {
		user, _ := services.User.Fetch(id)
		return user.Name
	}

	return f
}

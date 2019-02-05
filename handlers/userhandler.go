package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request){


	//check that user is logged in
		// code



	//fetch users information from server


	// TODO remove this with actual information
	type class struct {
		Code string
		Name string
		Link string
	}

	type user struct {
		Name string
		PrimaryEmail string
		SecondaryEmail string
		Classes []class
	}

	mobile := class{"IMT1337", "Mobile Development", "#"}
	app := class{"IMT4200", "Application Development", "#"}
	cloud := class{"IMT8008", "Cloud Technologies", "/"}
	www := class{"IMT0880", "WWW technology", "/"}

	classes := []class{
		mobile,
		app,
		cloud,
		www,
	}

	data := user{"Ola Nordmann", "olano@stud.ntnu.no", "", classes}
	// TODO : end here


	//parse information with template
	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("web/user.html")
	if err != nil{
		log.Fatal(err)
	}

	temp.Execute(w, data)

}

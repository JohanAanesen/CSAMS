package controller

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//CourseGET serves class page to users
func CourseGET(w http.ResponseWriter, r *http.Request) {

	var course model.Course

	if !session.GetUserFromSession(r).Authenticated {
		LoginGET(w, r)
		return
	}

	//check if request has an classID
	id := r.FormValue("id")
	if id == "" {
		//redirect to error pageinfo
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	courseID, err := strconv.Atoi(id)
	if err != nil {
		//redirect to error pageinfo
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	//get user
	user := session.GetUserFromSession(r)

	//get info from db
	course = db.GetCourse(courseID)

	//check if user is an participant of said class or a teacher
	participant := db.UserExistsInCourse(user.ID, courseID) || user.ID == course.Teacher
	if !participant {
		log.Println("user not participant of class")
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	//fmt.Println(course) //todo(johan) remove this

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	md := []byte(course.Description)
	description := github_flavored_markdown.Markdown(md) //todo sanitize markdown

	v := view.New(r)
	v.Name = "course"
	v.Vars["Course"] = course
	v.Vars["User"] = user
	v.Vars["Description"] = template.HTML(description)
	v.Render(w)
}

//CourseListGET serves class list page to users
func CourseListGET(w http.ResponseWriter, r *http.Request) {

	//check if request has an classID
	if r.Method == http.MethodGet {

		id := r.FormValue("id")

		if id == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Id is %s\n", id)
	}
	//check if user is an participant of said class or a teacher

	//get classlist from db

	//parse info to html template
	temp, err := template.ParseFiles("template/test.html")
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}

	temp.Execute(w, nil)
}

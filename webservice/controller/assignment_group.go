package controller

import (
	"fmt"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// AssignmentGroupCreateGET handles GET-requests @ /assignment/group/create
func AssignmentGroupCreateGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	services := service.NewServices(db.GetDB())

	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	v := view.New(r)
	v.Name = "assignment/group/create"

	v.Vars["Assignment"] = assignment

	v.Render(w)
}

// AssignmentGroupCreatePOST handles POST-requests @ /assignment/group/create
func AssignmentGroupCreatePOST(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	currentUser := session.GetUserFromSession(r)

	services := service.NewServices(db.GetDB())

	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	group := model.Group{}

	group.AssignmentID = assignment.ID
	group.Name = r.FormValue("group_name")
	group.Creator = currentUser.ID

	groupID, err := services.GroupService.Insert(group)
	if err != nil {
		log.Println("group service, insert", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = services.GroupService.AddUser(int(groupID), currentUser.ID)
	if err != nil {
		log.Println("group service, add user", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/assignment/%d", assignment.ID), http.StatusFound)
}

// AssignmentGroupJoinGET handles GET-requests @ /assignment/{aid}/join_group/{gid}
func AssignmentGroupJoinGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//currentUser := session.GetUserFromSession(r)

	services := service.NewServices(db.GetDB())

	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/assignment/%d", assignment.ID), http.StatusFound)
}

// AssignmentGroupLeaveGET handles GET-requests @ /assignment/{aid}/leave_group
func AssignmentGroupLeaveGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//currentUser := session.GetUserFromSession(r)

	services := service.NewServices(db.GetDB())

	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/assignment/%d", assignment.ID), http.StatusFound)
}

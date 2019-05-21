package controller

import (
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// NotificationGET marks notification as unactive and redirects user
func NotificationGET(w http.ResponseWriter, r *http.Request) {
	// Get current user
	currentUser := session.GetUserFromSession(r)
	// Get URL variables
	vars := mux.Vars(r)

	notificationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	//fetch notification //also doubles as check that the notification belongs to user
	notification, err := services.Notification.FetchNotificationForUser(currentUser.ID, notificationID)
	if err != nil {
		log.Println("services, notification, FetchNotificationForUSer", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//mark notification as read/unactive
	err = services.Notification.MarkAsRead(notification.ID)
	if err != nil {
		log.Println("notification MarkAsRead", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//all a-ok
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//redirect
	http.Redirect(w, r, notification.URL, http.StatusFound)
}

package main

import (
	"net/http"
	"os"
	"../../internal/handlers"
	dbcon "../../db"
)


func main() {

	dbcon.InitDB(os.Getenv("SQLDB")) //env var SQLDB username:password@tcp(127.0.0.1:3306)/dbname 127.0.0.1 if run locally like xampp


	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/class", handlers.ClassHandler)
	http.HandleFunc("/class/list", handlers.ClassListHandler)
	http.HandleFunc("/user", handlers.UserHandler)
	http.HandleFunc("/admin", handlers.AdminHandler)
	http.HandleFunc("/assignment", handlers.AssignmentHandler)
	http.HandleFunc("/assignment/peer", handlers.AssignmentPeerHandler)
	http.HandleFunc("/assignment/auto", handlers.AssignmentAutoHandler)


	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) //all files within /assets are served as static files

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)

}
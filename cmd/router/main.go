package main

import (
	dbcon "../../db"
	"../../internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)


func main() {

	dbcon.InitDB(os.Getenv("SQLDB")) //env var SQLDB username:password@tcp(127.0.0.1:3306)/dbname 127.0.0.1 if run locally like xampp

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.MainHandler)
	router.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	router.HandleFunc("/login", handlers.LoginRequest).Methods("POST")
	router.HandleFunc("/class", handlers.ClassHandler)
	router.HandleFunc("/class/list", handlers.ClassListHandler)
	router.HandleFunc("/user", handlers.UserHandler)
	router.HandleFunc("/admin", handlers.AdminHandler)
	router.HandleFunc("/assignment", handlers.AssignmentHandler)
	router.HandleFunc("/assignment/peer", handlers.AssignmentPeerHandler)
	router.HandleFunc("/assignment/auto", handlers.AssignmentAutoHandler)


	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) //all files within /assets are served as static files

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)

}
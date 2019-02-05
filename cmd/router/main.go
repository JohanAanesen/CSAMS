package main

import (
	"../../handlers"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/class", handlers.ClassHandler)
	http.HandleFunc("/class/list", handlers.ClassListHandler)
	http.HandleFunc("/user", handlers.UserHandler)
	http.HandleFunc("/admin", handlers.AdminHandler)
	http.HandleFunc("/assignment", handlers.AssignmentHandler)
	http.HandleFunc("/assignment/peer", handlers.AssignmentPeerHandler)
	http.HandleFunc("/assignment/auto", handlers.AssignmentAutoHandler)

	/*
	r.HandleFunc("/posts", articlesHandler).Methods("GET", "POST")
	r.HandleFunc("/posts/{id:[0-9]+}", articleHandler).Methods("GET")
	r.HandleFunc("/posts/delete", deleteArticleHandler).Methods("DELETE")
	*/

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) //all files within /static are served as static files

	log.Fatal(http.ListenAndServe(util.GetPort(), nil))
}
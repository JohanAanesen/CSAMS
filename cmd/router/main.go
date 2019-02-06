package main

import (
	dbcon "github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	dbcon.InitDB(os.Getenv("SQLDB")) //env var SQLDB username:password@tcp(127.0.0.1:3306)/dbname 127.0.0.1 if run locally like xampp

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.MainHandler).Methods("GET")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	router.HandleFunc("/login", handlers.LoginRequest).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterRequest).Methods("POST")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")
	router.HandleFunc("/class", handlers.ClassHandler).Methods("GET")
	router.HandleFunc("/class/list", handlers.ClassListHandler).Methods("GET")
	router.HandleFunc("/user", handlers.UserHandler).Methods("GET")
	router.HandleFunc("/user", handlers.UserUpdateRequest).Methods("POST")
	router.HandleFunc("/admin", handlers.AdminHandler).Methods("GET")
	router.HandleFunc("/assignment", handlers.AssignmentHandler).Methods("GET")
	router.HandleFunc("/assignment/peer", handlers.AssignmentPeerHandler).Methods("GET")
	router.HandleFunc("/assignment/auto", handlers.AssignmentAutoHandler).Methods("GET")

	//serve /assets as a static, always available folder
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}

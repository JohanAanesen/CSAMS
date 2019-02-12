package main

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/route"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/database"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/server"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"os"
)

func main() {
	var cfg = &config.Configuration{}
	cfg, _ = config.Load("config/config.json")

	// Configure Session
	session.Configure(cfg.Session)

	// Configure Database
	database.Configure(cfg.Database)

	db.InitDB(os.Getenv("SQLDB"))

	// Configure View
	view.Configure(cfg.View)
	view.LoadTemplate(cfg.Template.Root, cfg.Template.Children)

	// Run Server
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), cfg.Server)
}

package main

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/route"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/database"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/server"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"os"
)

func main() {
	var cfg = &config.Configuration{}
	cfg, _ = config.Load("config/config.json")

	database.Configure(cfg.Database)
	db.InitDB(os.Getenv("SQLDB"))

	view.Configure(cfg.View)
	view.LoadTemplate(cfg.Template.Root, cfg.Template.Children)

	server.Run(route.LoadHTTP(), route.LoadHTTPS(), cfg.Server)

	/*
		// Instantiate logger
		logger := log.New(os.Stdout, "main: ", log.LstdFlags | log.Lshortfile)

		// Instantiate database
		db := database.New("")

		// Instantiate config
		cfg := config.New()

		// Instantiate router
		router := mux.NewRouter().StrictSlash(true)

		// Setup routes
		route.Setup(router, logger, db)

		// Instantiate server
		srv := server.New(router, util.GetPort())
	*/

}

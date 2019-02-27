package main

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/route"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/server"
)

func main() {
	// Initialize config
	var cfg = config.Initialize("config/config.json")

	defer db.CloseDB()

	// Run Server
	server.Run(route.Load(), route.LoadHTTPS(), cfg.Server)
}

package main

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/route"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/server"
)

func main() {
	// Initialize config
	var cfg = config.Initialize()

	// Run Server
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), cfg.Server)
}

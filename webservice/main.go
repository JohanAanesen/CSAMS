package main

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/route"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/server"
)

func main() {
	// Initialize config
	var cfg = config.Initialize()
	fmt.Println("hurray started")
	defer db.CloseDB()
	fmt.Println("hurray started")
	// Run Server
	server.Run(route.Load(), route.LoadHTTPS(), cfg.Server)
}

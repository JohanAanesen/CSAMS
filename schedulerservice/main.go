package main

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/schedulerservice/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/schedulerservice/model"
)

func main() {
	// Initialize config
	var cfg = Initialize()

	// Initialize timers
	model.InitializeTimers()

	defer db.CloseDB()

	// Run Server
	Run(LoadHTTP(), LoadHTTPS(), cfg.Server)
}

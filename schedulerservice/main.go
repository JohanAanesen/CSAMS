package main

import (
	"github.com/JohanAanesen/CSAMS/schedulerservice/db"
	"github.com/JohanAanesen/CSAMS/schedulerservice/model"
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

package main

import (
	"github.com/JohanAanesen/CSAMS/webservice/route"
	"github.com/JohanAanesen/CSAMS/webservice/shared/config"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/server"
)

func main() {
	// Initialize config
	var cfg = config.Initialize("config/config.json")

	defer db.CloseDB()

	// Run Server
	server.Run(route.Load(), route.LoadHTTPS(), cfg.Server)
}

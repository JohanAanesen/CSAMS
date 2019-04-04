package main

import (
	"github.com/JohanAanesen/CSAMS/webservice/route"
	"github.com/JohanAanesen/CSAMS/webservice/shared/config"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/server"
	"log"
)

func main() {
	// Initialize config
	var cfg = config.Initialize("webservice/config/config.json")

	if cfg == nil {
		log.Fatal("could not load config, shutting down")
		return
	}

	defer db.CloseDB()

	// Run Server
	server.Run(route.Load(), route.LoadHTTPS(), cfg.Server)
}

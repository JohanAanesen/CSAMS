package main

import "fmt"

func main() {
	// Initialize config
	var cfg = Initialize()

	defer CloseDB()

	fmt.Println("hurray started")

	// Run Server
	Run(LoadHTTP(), LoadHTTPS(), cfg.Server)
}
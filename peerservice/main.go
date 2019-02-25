package main


func main() {
	// Initialize config
	var cfg = Initialize()

	defer CloseDB()

	// Run Server
	Run(LoadHTTP(), LoadHTTPS(), cfg.Server)
}
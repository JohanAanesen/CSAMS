package util

import "os"

//GetPort retrieves port from env var or sets 8080 default
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

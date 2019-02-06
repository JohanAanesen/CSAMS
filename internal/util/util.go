package util

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"log"
	"os"
)

// Retrieve and check if PORT is set in environment, if not, set it to 8080
// Returns port with a colon prefix
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

// Loads a JSON-file from disk, and decodes it to a Menu-struct
func LoadMenuConfig(file string) page.Menu {
	var menu page.Menu
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		log.Fatal(err)
		return page.Menu{}
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&menu)

	return menu
}

// Loads a JSON-file from disk, and decodes it to a Courses-struct
func LoadCoursesConfig(file string) page.Courses {
	var course page.Courses
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		log.Fatal(err)
		return page.Courses{}
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&course)

	return course
}


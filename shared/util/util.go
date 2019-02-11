package util

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"log"
	"os"
)

// GetPort gets environment PORT or sets it to 8080
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

// LoadMenuConfig loads a JSON-file from disk, and decodes it to a Menu-struct
func LoadMenuConfig(file string) model.Menu {
	var menu model.Menu
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		log.Fatal(err)
		return model.Menu{}
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&menu)

	return menu
}

// LoadCoursesConfig loads a JSON-file from disk, and decodes it to a Courses-struct
func LoadCoursesConfig(file string) model.Courses {
	var course model.Courses
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		log.Fatal(err)
		return model.Courses{}
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&course)

	return course
}

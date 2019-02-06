package util

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
	"log"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

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


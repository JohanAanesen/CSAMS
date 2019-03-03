package main

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/schedulerservice/db"
	"io/ioutil"
	"log"
	"os"
)

//Configuration struct
type Configuration struct {
	Database *db.MySQLInfo `json:"database"`
	Server   *Server    `json:"server"`
}

//LoadConfig loads configuration from config.json
func LoadConfig(configFile string) (*Configuration, error) {
	// Open file
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("could not open config file: %v\n", err)
		return &Configuration{}, err
	}

	// Close file
	defer file.Close()
	// Read file
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read all from file: %v\n", err)
		return &Configuration{}, err
	}

	// Declare Configuration
	cfg := Configuration{}
	// Unmarshal JSON to Configuration-object
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatalf("could not unmarshal json: %v\n", err)
		return &Configuration{}, err
	}

	return &cfg, nil
}

//Initialize loads the config.json file and sets up database
func Initialize() *Configuration {
	var cfg = &Configuration{}
	cfg, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	// Configure Database
	db.ConfigureDB(cfg.Database)
	db.OpenDB()

	return cfg
}

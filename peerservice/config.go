package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Database *MySQLInfo  `json:"database"`
	Server   *Server `json:"server"`
}

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

func Initialize() *Configuration {
	var cfg = &Configuration{}
	cfg, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	// Configure Database
	ConfigureDB(cfg.Database)
	OpenDB()

	return cfg
}

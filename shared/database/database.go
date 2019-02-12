package database

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db *sql.DB
)

// MySQLInfo struct
type MySQLInfo struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	ParseTime bool `json:"parseTime"`
}

func Configure(info *MySQLInfo) {
	// root:@tcp(127.0.0.1:3306)/cs53
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t", info.Username, info.Password, info.Hostname, info.Port, info.Database, info.ParseTime))
	if err != nil {
		panic(err.Error())
	}

	if err = database.Ping(); err != nil {
		panic(err.Error())
	}

	db = database
	log.Print("database connected\n")
}

func Get() *sql.DB {
	return db
}


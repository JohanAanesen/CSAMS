package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //database driver
)

const driverName = "mysql"

var (
	dataSourceName string
	db             *sql.DB
)

// MySQLInfo struct
type MySQLInfo struct {
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	ParseTime bool   `json:"parseTime"`
}

//ConfigureDB loads database connection string
func ConfigureDB(info *MySQLInfo) {
	// root:@tcp(127.0.0.1:3306)/cs53
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t", info.Username, info.Password, info.Hostname, info.Port, info.Database, info.ParseTime)
}

//OpenDB creates connection with database
func OpenDB() {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
}

//GetDB returns db object
func GetDB() *sql.DB {
	return db
}

//CloseDB closes db
func CloseDB() {
	db.Close()
}

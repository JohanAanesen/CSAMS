package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //database driver
)

const driverName = "mysql"

var (
	dataSourceName string
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
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t", info.Username, info.Password, info.Hostname, info.Port, info.Database, info.ParseTime)
}

func Open() {
	db, _ = sql.Open(driverName, dataSourceName)
}

func Get() *sql.DB {
	return db
}

func Close() {
	db.Close()
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args)
}

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //database driver
	"log"
	"os"
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
// TODO (Svein): Is DB needed at the end? We are in db-package, usage: db.Configure() vs db.ConfigureDB)
func ConfigureDB(info *MySQLInfo) {
	// root:@tcp(127.0.0.1:3306)/cs53
	if os.Getenv("DATABASE_URL") == "" {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t", info.Username, info.Password, info.Hostname, info.Port, info.Database, info.ParseTime)
	} else {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=%t",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_URL"),
			os.Getenv("DATABASE_DATABASE"),
			info.ParseTime)
	}
}

//OpenDB creates connection with database
// TODO (Svein): Is DB needed at the end? We are in db-package, usage: db.Open() vs db.OpenDB()
func OpenDB() {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("could not open sql connection, error:", err)
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println("could not ping database, error:", err)
		panic(err.Error())
	}
}

//GetDB returns db object
// TODO (Svein): Is DB needed at the end? We are in db-package, usage: db.Get() vs db.GetDB()
func GetDB() *sql.DB {
	return db
}

//CloseDB closes db
// TODO (Svein): Is DB needed at the end? We are in db-package, usage: db.Close() vs db.CloseDB()
func CloseDB() {
	db.Close()
}

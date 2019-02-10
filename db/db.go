package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //database driver
	"github.com/gorilla/sessions"
	"os"
)

//DB global DB connection variable
var DB *sql.DB

//CookieStore global var for session management
var CookieStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//InitDB initializes the database
func InitDB(dataSourceName string) {
	var err error

	DB, err = sql.Open("mysql", dataSourceName)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}

}


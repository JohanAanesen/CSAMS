package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"os"
)

var DB *sql.DB
var CookieStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))


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

func UserAuth(name string, password string)bool{

	rows, _ := DB.Query("SELECT name FROM users WHERE name = ? AND password = ?", name, password)

	for rows.Next(){ //todo all of this lol
		 var Name string

		 rows.Scan(&Name)

		 if Name != ""{
		 	return true
		 }
	}

	return false
}

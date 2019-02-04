package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
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

func UserAuth(email string, password string)bool{

	rows, _ := DB.Query("SELECT password FROM users WHERE email = ?", email)

	for rows.Next(){ //todo all of this lol
		 var hash string

		 rows.Scan(&hash)

		 if CheckPasswordHash(password, hash){
		 	fmt.Println("all guchi")
		 	return true
		 }
	}

	return false
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
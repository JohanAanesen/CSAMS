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

func UserAuth(email string, password string) (int, bool) {

	rows, err := DB.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return 0, false
	}

	for rows.Next() {
		var id int
		var hash string

		rows.Scan(&id, &hash)

		if CheckPasswordHash(password, hash) {
			fmt.Println("all guchi")
			return id, true
		}
	}

	defer rows.Close()

	return 0, false
}

func RegisterUser(email string, password string) (int, bool) {
	pass, err := HashPassword(password)
	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return 0, false
	}

	rows, err := DB.Query("INSERT INTO users(email, teacher, password) VALUES(?, 0, ?)", email, pass)
	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return 0, false
	}

	defer rows.Close()

	return UserAuth(email, password) //fetch userid through existing method
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

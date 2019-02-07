package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //database driver
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
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

//UserAuth authenticates users
func UserAuth(email string, password string) (int, string, bool) {

	rows, err := DB.Query("SELECT id, name, password FROM users WHERE email_student = ?", email)
	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return 0, "", false
	}

	for rows.Next() {
		var id int
		var hash string
		var name string

		rows.Scan(&id, &name, &hash)

		if CheckPasswordHash(password, hash) {
			return id, name, true
		}
	}

	defer rows.Close()

	return 0, "", false
}

//RegisterUser registers users to database
func RegisterUser(name string, email string, password string) (int, string, bool) {
	pass, err := hashPassword(password)
	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return 0, "", false
	}

	rows, err := DB.Query("INSERT INTO users(name, email_student, teacher, password) VALUES(?, ?, 0, ?)", name, email, pass)
	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return 0, "", false
	}

	defer rows.Close()

	return UserAuth(email, password) //fetch user-id through existing method
}

// TODO : fix this function >:(
// GetUSer returns the email_private and teacher-id
func GetUser(userID int) (int, string, string, int, string, string) {

	rows, err := DB.Query("SELECT id, name, email_student, teacher, email_private, password FROM users WHERE id = ?", userID)
	if err != nil {
		// TODO : log error
		fmt.Println(err.Error())
		return -1, "", "", -1, "", ""
	}

	for rows.Next() {
		var id int
		var name string
		var emailStudent sql.NullString // Johan fixed: https://golang.org/pkg/database/sql/#NullString
		var teacher int
		var emailPrivate string
		var password string // this is not showing

		rows.Scan(&id, &name, &emailStudent, &teacher, &emailPrivate, &password)

		return id, name, emailStudent.String, teacher, emailPrivate, password
	}

	defer rows.Close()

	return -1, "", "", -1, "", ""
}

// TODO : remove this function and fix the GetUser function >:(
func GetHash(id int) string{
	rows, err := DB.Query("SELECT password FROM users WHERE id = ?", id)
	if err != nil {
		// TODO : log error
		fmt.Println(err.Error())
		return ""
	}

	for rows.Next() {
		var password string

		rows.Scan(&password)

		return password
	}

	defer rows.Close()

	return ""
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

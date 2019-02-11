package db

import (
	"database/sql"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
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
func UserAuth(email string, password string) (model.User, bool) {
	rows, err := DB.Query("SELECT id, password FROM users WHERE email_student = ?", email)

	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return model.User{Authenticated: false}, false
	}

	for rows.Next() {
		var id int
		var hash string

		rows.Scan(&id, &hash)
		if err != nil {
			//todo log error
			fmt.Println(err.Error())
			return model.User{Authenticated: false}, false
		}

		if checkPasswordHash(password, hash) {
			return GetUser(id), true
		}
	}

	defer rows.Close()

	return model.User{Authenticated: false}, false
}

//RegisterUser registers users to database
func RegisterUser(name string, email string, password string) (model.User, bool) {
	pass, err := hashPassword(password)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return model.User{Authenticated: false}, false
	}

	rows, err := DB.Query("INSERT INTO users(name, email_student, teacher, password) VALUES(?, ?, 0, ?)", name, email, pass)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return model.User{Authenticated: false}, false
	}

	defer rows.Close()

	return UserAuth(email, password) //fetch userid through existing method
}

func GetUser(userID int) model.User {
	rows, err := DB.Query("SELECT id, name, email_student, email_private, teacher FROM users WHERE id = ?", userID)
	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return model.User{Authenticated: false}
	}

	for rows.Next() {
		var user model.User
		var id int
		var name string
		var emailStudent string
		var emailPrivate sql.NullString
		var teacher bool

		err := rows.Scan(&id, &name, &emailStudent, &emailPrivate, &teacher)
		if err != nil {
			//todo log error
			fmt.Println(err.Error())
			return model.User{Authenticated: false}
		}

		user.ID = userID
		user.Name = name
		user.EmailStudent = emailStudent
		user.EmailPrivate = emailPrivate.String
		user.Teacher = teacher

		return user
	}

	defer rows.Close()

	return model.User{Authenticated: false}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

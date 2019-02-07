package db

import (
	"database/sql"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/structs"
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

	rows, err := DB.Query("SELECT * FROM users WHERE id = ?", userID)
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

// GetCourseToUser returns all the courses to the user
func GetCoursesToUser(userID int) []structs.CourseDB {

	// Create an empty courses array
	var courses []structs.CourseDB

	rows, err := DB.Query("SELECT course.* FROM course INNER JOIN usercourse ON course.id = usercourse.courseid WHERE usercourse.userid = ?", userID)
	if err != nil {
		fmt.Println(err.Error()) // TODO : log error

		// returns empty course array if it fails
		return courses
	}

	for rows.Next() {
		var id int
		var courseCode string
		var courseName string
		var teacher int
		var info string
		var link1 string
		var link2 string
		var link3 string

		rows.Scan(&id, &courseCode, &courseName, &teacher, &info, &link1, &link2, &link3)

		// Add course to courses array
		courses = append(courses, structs.CourseDB{
			Id:         id,
			CourseCode: courseCode,
			CourseName: courseName,
			Teacher:    teacher,
			Info:       info,
			Link1:      link1,
			Link2:      link2,
			Link3:      link3,
		})
	}

	return courses
}

// TODO : remove this function and fix the GetUser function >:(
// GetHash returns the users hashed password
func GetHash(id int) string {
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

// UpdateUserName updates the users name in the db
func UpdateUserName(userID int, newName string) bool {

	rows, err := DB.Query("UPDATE users SET name = ? WHERE id = ?", newName, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	} else {
		// TODO : maybe add confirmation or something
		defer rows.Close()
		return true
	}

}

// UpdateUserEmail updates the users email in the db
func UpdateUserEmail(userID int, email string) bool {
	rows, err := DB.Query("UPDATE users SET email_private = ? WHERE id = ?", email, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	} else {
		// TODO : maybe add confirmation or something
		defer rows.Close()
		return true
	}
}

// UpdateUserPassword updates the users password in the db
func UpdateUserPassword(userID int, password string) bool {

	// Hash the password first
	pass, err := hashPassword(password)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	}

	rows, err := DB.Query("UPDATE users SET password = ? WHERE id = ?", pass, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	} else {
		// TODO : maybe add confirmation or something
		defer rows.Close()
		return true
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

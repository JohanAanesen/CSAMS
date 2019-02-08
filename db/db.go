package db

import (
	"database/sql"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/page"
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

		if CheckPasswordHash(password, hash) {
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

	return UserAuth(email, password) //fetch user-id through existing method
}

// GetCourseToUser returns all the courses to the user
func GetCoursesToUser(userID int) page.Courses {

	// Create an empty courses array
	var courses page.Courses

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
		var description string
		var year string
		var semester string

		rows.Scan(&id, &courseCode, &courseName, &teacher, &description, &year, &semester)

		// Add course to courses array
		courses.Items = append(courses.Items, page.Course{
			Id:          id,
			Code:  courseCode,
			Name:  courseName,
			Teacher:     teacher,
			Description: description,
			Year:        year,
			Semester:    semester,
		})
	}

	return courses
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

func GetUser(userID int) (model.User) {
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

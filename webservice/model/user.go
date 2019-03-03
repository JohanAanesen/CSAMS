package model

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Users hold the data for a slice of Course-struct
type Users struct {
	Items []User `json:"users"`
}

//User struct to hold session data
type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	EmailStudent  string `json:"emailstudent"`
	EmailPrivate  string `json:"emailprivate"`
	Teacher       bool   `json:"teacher"`
	Authenticated bool   `json:"authenticated"`
}

// UpdateUserName updates the users name in the db
func UpdateUserName(userID int, newName string) bool {

	rows, err := db.GetDB().Query("UPDATE users SET name = ? WHERE id = ?", newName, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	}

	defer rows.Close()

	return true
}

// UserIsRreviewer Checks if the user (userID) can review another user(reviewUserID)
func UserIsRreviewer(UserID int, assignmentID int, submissionID int64, reviewUserID int) bool {

	// Run query
	rows, err := db.GetDB().Query("SELECT * from peer_reviews WHERE assignment_id = ? AND submission_id = ? AND review_user_id = ? AND user_id = ?", assignmentID, submissionID, reviewUserID, UserID)
	if err != nil {
		// Return false if user can not review
		return false
	}

	// If there was a match
	if rows.Next() {
		return true
	}

	defer rows.Close()

	// Return true if user can review
	return false
}

//UpdateUserEmail updates the users email in the db
func UpdateUserEmail(userID int, email string) bool {
	rows, err := db.GetDB().Query("UPDATE users SET email_private = ? WHERE id = ?", email, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	}

	defer rows.Close()

	return true
}

//UpdateUserPassword updates the users password in the db
func UpdateUserPassword(userID int, password string) bool {

	// Hash the password first
	pass, err := hashPassword(password)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	}

	rows, err := db.GetDB().Query("UPDATE users SET password = ? WHERE id = ?", pass, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	}

	defer rows.Close()

	return true
}

//GetUser retrieves an user from DB through userID
func GetUser(userID int) User {
	rows, err := db.GetDB().Query("SELECT id, name, email_student, email_private, teacher FROM users WHERE id = ?", userID)
	if err != nil {
		log.Println(err.Error())
		return User{Authenticated: false}
	}

	for rows.Next() {
		var id int
		var name string
		var emailStudent string
		var emailPrivate sql.NullString
		var teacher bool

		err := rows.Scan(&id, &name, &emailStudent, &emailPrivate, &teacher)
		if err != nil {
			log.Println(err.Error())
			return User{Authenticated: false}
		}

		// Return the user in a User struct from model folder
		return User{
			ID:            userID,
			Name:          name,
			EmailStudent:  emailStudent,
			EmailPrivate:  emailPrivate.String,
			Teacher:       teacher,
			Authenticated: true,
		}
	}

	defer rows.Close()

	return User{Authenticated: false}
}

// GetHash returns the users hashed password
func GetHash(id int) string {
	rows, err := db.GetDB().Query("SELECT password FROM users WHERE id = ?", id)
	if err != nil {
		log.Println(err.Error())
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

// GetUsersToCourse returns all users to one course
func GetUsersToCourse(courseID int) Users {

	//Create an empty courses array
	var users Users

	rows, err := db.GetDB().Query("SELECT users.id, users.name, users.email_student, teacher FROM users INNER JOIN usercourse ON users.id = usercourse.userid WHERE usercourse.courseid = ?", courseID)
	if err != nil {
		log.Println(err.Error())

		// returns empty course array if it fails
		return users
	}

	for rows.Next() {
		var id int
		var name string
		var email string
		var teacher bool

		rows.Scan(&id, &name, &email, &teacher)

		// Add course to courses array
		users.Items = append(users.Items, User{
			ID:           id,
			Name:         name,
			EmailStudent: email,
			Teacher:      teacher,
		})
	}

	return users
}

//UserAuth authenticates users
func UserAuth(email string, password string) (User, bool) {
	rows, err := db.GetDB().Query("SELECT id, password FROM users WHERE email_student = ?", email)

	if err != nil {
		log.Println(err.Error())
		return User{Authenticated: false}, false
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var hash string

		rows.Scan(&id, &hash)

		if err != nil {
			log.Println(err.Error())
			return User{Authenticated: false}, false
		}

		if CheckPasswordHash(password, hash) {
			return GetUser(id), true
		}
	}

	return User{Authenticated: false}, false
}

//RegisterUser registers users to database
func RegisterUser(name string, email string, password string) (User, bool) {
	pass, err := hashPassword(password)

	if err != nil {
		log.Fatal(err.Error())
		return User{Authenticated: false}, false
	}

	_, err = db.GetDB().Exec("INSERT INTO users(name, email_student, teacher, password) VALUES(?, ?, 0, ?)", name, email, pass)

	if err != nil {
		log.Fatal(err.Error())
		return User{Authenticated: false}, false
	}

	return UserAuth(email, password) //fetch user-id through existing method
}

//CheckPasswordHash compares password string and hashed string
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

package db

import (
	"database/sql"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"log"
)

// TODO (Svein): Move this to model/user.go

// UpdateUserName updates the users name in the db
func UpdateUserName(userID int, newName string) bool {

	rows, err := GetDB().Query("UPDATE users SET name = ? WHERE id = ?", newName, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	}

	defer rows.Close()

	return true
}

//UpdateUserEmail updates the users email in the db
func UpdateUserEmail(userID int, email string) bool {
	rows, err := GetDB().Query("UPDATE users SET email_private = ? WHERE id = ?", email, userID)

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

	rows, err := GetDB().Query("UPDATE users SET password = ? WHERE id = ?", pass, userID)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return false
	}

	defer rows.Close()

	return true
}

//GetUser retrieves an user from DB through userID
func GetUser(userID int) model.User {
	rows, err := GetDB().Query("SELECT id, name, email_student, email_private, teacher FROM users WHERE id = ?", userID)
	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return model.User{Authenticated: false}
	}

	for rows.Next() {
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

		// Return the user in a User struct from model folder
		return model.User{
			ID:           userID,
			Name:         name,
			EmailStudent: emailStudent,
			EmailPrivate: emailPrivate.String,
			Teacher:      teacher,
		}
	}

	defer rows.Close()

	return model.User{Authenticated: false}
}

// GetHash returns the users hashed password
func GetHash(id int) string {
	rows, err := GetDB().Query("SELECT password FROM users WHERE id = ?", id)
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

func IsParticipant(userID int, courseID int) bool{
	particpant := false

	rows, err := GetDB().Query("SELECT userid, courseid FROM usercourse WHERE userid = ? AND courseid = ?", userID, courseID)
	if err != nil {
		// TODO : log error
		fmt.Println(err.Error())
		return particpant
	}

	for rows.Next() {
		var user int
		var course int
		rows.Scan(&user, &course)

		if user == userID && course == courseID{
			particpant = true
		}
	}

	defer rows.Close()

	return particpant
}

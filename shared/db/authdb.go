package db

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

//UserAuth authenticates users
func UserAuth(email string, password string) (model.User, bool) {
	rows, err := GetDB().Query("SELECT id, password FROM users WHERE email_student = ?", email)

	if err != nil {
		//todo log error
		fmt.Println(err.Error())
		return model.User{Authenticated: false}, false
	}

	defer rows.Close()

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

	rows, err := GetDB().Query("INSERT INTO users(name, email_student, teacher, password) VALUES(?, ?, 0, ?)", name, email, pass)

	if err != nil {
		//todo log error
		log.Fatal(err.Error())
		return model.User{Authenticated: false}, false
	}

	defer rows.Close()

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

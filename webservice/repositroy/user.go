package repositroy

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
)

// UserRepository struct
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository func
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Fetch func
func (repo *UserRepository) Fetch(id int) (*model.User, error) {
	result := model.User{}
	result.Authenticated = false

	query := "SELECT id, name, email_student, teacher, email_private FROM users WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.EmailStudent,
			&result.Teacher, &result.EmailPrivate)
		if err != nil {
			return &result, err
		}
	}

	result.Authenticated = true
	return &result, err
}

// FetchHash func
func (repo *UserRepository) FetchHash(id int) (string, error) {
	var result string

	query := "SELECT password FROM users WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return result, err
		}
	}

	return result, err
}

// FetchAll func
func (repo *UserRepository) FetchAll() ([]*model.User, error) {
	result := make([]*model.User, 0)

	query := "SELECT id, name, email_student, teacher, email_private FROM users"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.User{}

		err = rows.Scan(&temp.ID, &temp.Name, &temp.EmailStudent,
			&temp.Teacher, &temp.EmailPrivate)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchAllFromCourse fetches all course participants in sorted list by teacher, name
func (repo *UserRepository) FetchAllFromCourse(courseID int) ([]*model.User, error) {
	result := make([]*model.User, 0)

	query := "SELECT u.id, u.name, u.email_student, u.teacher, u.email_private FROM users AS u INNER JOIN usercourse AS uc ON u.id = uc.userid WHERE uc.courseid = ? ORDER BY u.teacher DESC, u.name ASC"

	rows, err := repo.db.Query(query, courseID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.User{}

		err = rows.Scan(&temp.ID, &temp.Name, &temp.EmailStudent,
			&temp.Teacher, &temp.EmailPrivate)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchAllStudentsFromCourse func
func (repo *UserRepository) FetchAllStudentsFromCourse(courseID int) ([]*model.User, error) {
	result := make([]*model.User, 0)

	query := "SELECT u.id, u.name, u.email_student, u.teacher, u.email_private FROM users AS u INNER JOIN usercourse AS uc ON u.id = uc.userid WHERE uc.courseid = ? AND u.teacher = 0"

	rows, err := repo.db.Query(query, courseID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.User{}

		err = rows.Scan(&temp.ID, &temp.Name, &temp.EmailStudent,
			&temp.Teacher, &temp.EmailPrivate)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// EmailExists Checks if a user with said email exists and returns userid
func (repo *UserRepository) EmailExists(email string) (bool, int, error) {

	query := "SELECT id FROM users WHERE email_student = ? OR email_private = ?"

	rows, err := repo.db.Query(query, email, email)
	if err != nil {
		return false, -1, err
	}

	defer rows.Close()

	// If rows.next is true, there were a match
	if rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return false, -1, err
		}

		return true, id, nil
	}

	return false, -1, err
}

// Insert func
func (repo *UserRepository) Insert(user model.User, password string) (int, error) {
	var id int64

	query := "INSERT INTO users (name, email_student, teacher, password) VALUES (?, ?, 0, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query, user.Name, user.EmailStudent, password)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), err
}

// Update func
func (repo *UserRepository) Update(id int, user model.User) error {
	if id != user.ID {
		return errors.New("update user, id mismatch")
	}

	query := "UPDATE users SET name = ?, email_private = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, user.Name, user.EmailPrivate, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

// UpdatePassword func
func (repo *UserRepository) UpdatePassword(id int, hashedPassword string) error {
	query := "UPDATE users SET password = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, hashedPassword, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

// Delete func
func (repo *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

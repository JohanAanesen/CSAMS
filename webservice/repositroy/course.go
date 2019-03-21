package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
)

// CourseRepository struct
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository func
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Fetch func
func (repo *CourseRepository) Fetch(id int) (*model.Course, error) {
	result := model.Course{}

	query := "SELECT id, hash, coursecode, coursename, teacher, description, year, semester FROM course WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Hash, &result.Code,
			&result.Name, &result.Teacher, &result.Description,
			&result.Year, &result.Semester)
		if err != nil {
			return &result, err
		}
	}

	return &result, err
}

// FetchAll func
func (repo *CourseRepository) FetchAll() ([]*model.Course, error) {
	result := make([]*model.Course, 0)

	query := "SELECT id, hash, coursecode, coursename, teacher, description, year, semester FROM course"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.Course{}

		err = rows.Scan(&temp.ID, &temp.Hash, &temp.Code,
			&temp.Name, &temp.Teacher, &temp.Description,
			&temp.Year, &temp.Semester)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}
// FetchAllForUserOrdered func
func (repo *CourseRepository) FetchAllForUserOrdered(userID int) ([]*model.Course, error) {
	result := make([]*model.Course, 0)

	query := "SELECT c.id, c.hash, c.coursecode, c.coursename, c.teacher, c.description, c.year, c.semester FROM course AS c INNER JOIN usercourse AS uc ON c.id = uc.courseid WHERE uc.userid = ? ORDER BY c.year DESC, c.semester ASC, c.coursename ASC"

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.Course{}

		err = rows.Scan(&temp.ID, &temp.Hash, &temp.Code,
			&temp.Name, &temp.Teacher, &temp.Description,
			&temp.Year, &temp.Semester)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}


// Insert func
func (repo *CourseRepository) Insert(course model.Course) (int, error) {
	var id int64

	query := "INSERT INTO course (hash, coursecode, coursename, teacher, description, year, semester) VALUES (?, ?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query, course.Hash, course.Code,
		course.Name, course.Teacher, course.Description,
		course.Year, course.Semester)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), err
}

// InsertUser func
func (repo *CourseRepository) InsertUser(userID, courseID int) error {
	query := "INSERT INTO usercourse (userid, courseid) VALUES (?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, userID, courseID)
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

// UserInCourse func
func (repo *CourseRepository) UserInCourse(userID, courseID int) (bool, error) {
	query := "SELECT courseid FROM usercourse WHERE userid = ? AND courseid = ?"

	rows, err := repo.db.Query(query, userID, courseID)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp int

		err = rows.Scan(&temp)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, err
}

// Update func
func (repo *CourseRepository) Update(id int, course model.Course) error {
	query := "UPDATE course SET coursecode = ?, coursename = ?, description = ?, semester = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, course.Code, course.Name, course.Description, course.Semester, id)
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
func (repo *CourseRepository) Delete(id int) error {
	query := "DELETE FROM course WHERE id = ?"

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

// RemoveUser func
func (repo *CourseRepository) RemoveUser(userID, courseID int) error {
	query := "DELETE FROM usercourse WHERE userid = ? AND courseid = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, userID, courseID)
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

// FetchAllForUser func
func (repo *CourseRepository) FetchAllForUser(userID int) ([]*model.Course, error) {
	result := make([]*model.Course, 0)

	query := "SELECT c.id, c.hash, c.coursecode, c.coursename, c.teacher, c.description, c.year, c.semester FROM course AS c INNER JOIN usercourse AS uc ON c.id = uc.courseid WHERE uc.userid = ? ORDER BY c.year DESC, c.semester ASC, c.coursename ASC"

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.Course{}

		err = rows.Scan(&temp.ID, &temp.Hash, &temp.Code,
			&temp.Name, &temp.Teacher, &temp.Description,
			&temp.Year, &temp.Semester)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

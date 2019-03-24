package service

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
)

var (
	// ErrUserAlreadyInCourse error
	ErrUserAlreadyInCourse = errors.New("user already in course")
)

// CourseService struct
type CourseService struct {
	courseRepo *repositroy.CourseRepository
}

// NewCourseService func
func NewCourseService(db *sql.DB) *CourseService {
	return &CourseService{
		courseRepo: repositroy.NewCourseRepository(db),
	}
}

// Fetch func
func (s *CourseService) Fetch(id int) (*model.Course, error) {
	return s.courseRepo.Fetch(id)
}

// FetchAll func
func (s *CourseService) FetchAll() ([]*model.Course, error) {
	return s.courseRepo.FetchAll()
}

// FetchAllForUser func
func (s *CourseService) FetchAllForUser(userID int) ([]*model.Course, error) {
	return s.courseRepo.FetchAllForUser(userID)
}

// FetchAllForUserOrdered func
func (s *CourseService) FetchAllForUserOrdered(userID int) ([]*model.Course, error) {
	return s.courseRepo.FetchAllForUserOrdered(userID)
}

// Exists func
func (s *CourseService) Exists(hash string) *model.Course {
	result := model.Course{
		ID: -1,
	}

	courses, err := s.courseRepo.FetchAll()
	if err != nil {
		return &result
	}

	for _, course := range courses {
		if course.Hash == hash {
			return course
		}
	}

	return &result
}

// UserInCourse checks if user is in given course
func (s *CourseService) UserInCourse(userID, courseID int) (bool, error) {
	return s.courseRepo.UserInCourse(userID, courseID)
}

// AddUser to a single course
func (s *CourseService) AddUser(userID, courseID int) error {
	exists, err := s.UserInCourse(userID, courseID)
	if err != nil {
		return err
	}

	if exists {
		return ErrUserAlreadyInCourse
	}

	return s.courseRepo.InsertUser(userID, courseID)
}

// RemoveUser from a course
func (s *CourseService) RemoveUser(userID, courseID int) error {
	return s.courseRepo.RemoveUser(userID, courseID)
}

// Insert course into the database
func (s *CourseService) Insert(course model.Course) (int, error) {
	return s.courseRepo.Insert(course)
}

// Update a course in the database
func (s *CourseService) Update(id int, course model.Course) error {
	return s.courseRepo.Update(id, course)
}

// Delete a course in the database
func (s *CourseService) Delete(id int) error {
	return s.courseRepo.Delete(id)
}

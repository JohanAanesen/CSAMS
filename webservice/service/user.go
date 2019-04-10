package service

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
)

// UserService struct
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService func
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(db),
	}
}

// Fetch func
func (s *UserService) Fetch(id int) (*model.User, error) {
	return s.userRepo.Fetch(id)
}

// FetchHash func
func (s *UserService) FetchHash(id int) (string, error) {
	return s.userRepo.FetchHash(id)
}

// FetchAll func
func (s *UserService) FetchAll() ([]*model.User, error) {
	return s.userRepo.FetchAll()
}

// FetchAllFromCourse func
func (s *UserService) FetchAllFromCourse(courseID int) ([]*model.User, error) {
	return s.userRepo.FetchAllFromCourse(courseID)
}

// FetchAllStudentsFromCourse func
func (s *UserService) FetchAllStudentsFromCourse(courseID int) ([]*model.User, error) {
	return s.userRepo.FetchAllStudentsFromCourse(courseID)
}

// EmailExists checks if the email exists in emailprivate and emailstudent
func (s *UserService) EmailExists(email string) (bool, int, error) {
	return s.userRepo.EmailExists(email)
}

// RegisterWithHashing func
func (s *UserService) RegisterWithHashing(user model.User, password string) (int, error) {
	hashed, err := util.GenerateFromPassword(password)
	if err != nil {
		return 0, err
	}

	return s.userRepo.Insert(user, hashed)
}

// Register func
func (s *UserService) Register(user model.User, hashedPass string) (int, error) {
	return s.userRepo.Insert(user, hashedPass)
}

// Authenticate func
func (s *UserService) Authenticate(email, password string) (model.User, error) {
	// Create empty user
	result := model.User{}
	// Get all users
	users, err := s.userRepo.FetchAll()
	if err != nil {
		return result, err
	}
	// Set found to false
	found := false
	// Loop through all users
	for _, u := range users {
		// Check if any if it's emails match
		if u.EmailStudent == email || u.EmailPrivate.String == email {
			// Set user, and found to true
			result = *u
			found = true
			break
		}
	}
	// Not found
	if !found {
		return result, errors.New("user not found by given email")
	}
	// Get hash for user
	hash, err := s.FetchHash(result.ID)
	if err != nil {
		return result, err
	}
	// Compare hash and password
	err = util.CompareHashAndPassword(password, hash)
	if err != nil {
		return result, err
	}

	return result, err
}

// Update func
func (s *UserService) Update(id int, user model.User) error {
	return s.userRepo.Update(id, user)
}

// UpdatePassword func
func (s *UserService) UpdatePassword(id int, password string) error {
	hash, err := util.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(id, hash)
}

// FetchAllStudentEmails Fetches all student emails, primary default or secondary if not null
func (s *UserService) FetchAllStudentEmails(courseID int) ([]string, error) {

	// Create empty string array
	var result []string

	// Get all users from course
	users, err := s.userRepo.FetchAllFromCourse(courseID)
	if err != nil {
		return result, err
	}

	// Loop through all users
	for _, user := range users {

		// only append to array if user is an student
		if !user.Teacher {

			// Check if user has secondary/private email and append if yes
			if user.EmailPrivate.Valid {
				result = append(result, user.EmailPrivate.String)
			} else {
				// Append primary/student email
				result = append(result, user.EmailStudent)
			}
		}
	}

	return result, nil
}

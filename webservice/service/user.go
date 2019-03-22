package service

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
)

// UserService struct
type UserService struct {
	userRepo *repositroy.UserRepository
}

// NewUserService func
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		userRepo: repositroy.NewUserRepository(db),
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

// Register func
func (s *UserService) Register(user model.User, password string) (int, error) {
	hashed, err := util.GenerateFromPassword(password)
	if err != nil {
		return 0, err
	}

	exists, err := s.userRepo.EmailExists(user)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("email already exists")
	}

	return s.userRepo.Insert(user, hashed)
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

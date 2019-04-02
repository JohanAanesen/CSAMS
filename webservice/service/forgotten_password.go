package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
)

// ForgottenPassService struct
type ForgottenPassService struct {
	forgottenPassRepo *repositroy.ForgottenPassRepository
}

// NewForgottenPassService func
func NewForgottenPassService(db *sql.DB) *ForgottenPassService {
	return &ForgottenPassService{
		forgottenPassRepo: repositroy.NewForgottenPassRepository(db),
	}
}

// Insert func
func (s *ForgottenPassService) Insert(forgottenPass model.ForgottenPass) (int, error) {
	var err error

	// First hash the hash for more security
	forgottenPass.Hash, err = util.GenerateFromPassword(forgottenPass.Hash)
	if err != nil {
		return -1, err
	}

	return s.forgottenPassRepo.Insert(forgottenPass)
}

// Match first hashes the hash and then checks if the hash match in the db
func (s *ForgottenPassService) Match(hash string) (bool, error) {

	//  First hash the hash for more security
	hashed, err := util.GenerateFromPassword(hash)
	if err != nil {
		return false, err
	}

	return s.forgottenPassRepo.Match(hashed)
}

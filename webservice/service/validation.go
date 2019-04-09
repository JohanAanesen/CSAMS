package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
)

// ValidationService struct
type ValidationService struct {
	validationRepo *repository.ValidationRepository
}

// NewValidationService func
func NewValidationService(db *sql.DB) *ValidationService {
	return &ValidationService{
		validationRepo: repository.NewValidationRepository(db),
	}
}

// Insert func
func (s *ValidationService) Insert(forgottenPass model.Validation) (int, error) {
	var err error

	// First hash the hash for more security
	forgottenPass.Hash, err = util.GenerateFromPassword(forgottenPass.Hash)
	if err != nil {
		return -1, err
	}

	return s.validationRepo.Insert(forgottenPass)
}

// Match first hashes the hash and then checks if the hash match in the db
func (s *ValidationService) Match(hash string) (bool, model.Validation, error) {

	// Fetch all rows
	forgottenPasses, err := s.validationRepo.FetchAll()
	if err != nil {
		return false, model.Validation{}, err
	}

	for _, item := range forgottenPasses {
		err := util.CompareHashAndPassword(hash, item.Hash)

		// If there is a match
		if err == nil {
			return true, *item, err
		}
	}

	return false, model.Validation{}, nil
}

// UpdateValidation updates the validation to a link
func (s *ValidationService) UpdateValidation(id int, state bool) error {
	return s.validationRepo.UpdateValidation(id, state)
}

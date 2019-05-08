package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
)

// UserPendingService struct
type UserPendingService struct {
	userPendingRepo *repository.UserPendingRepository
}

// NewUserPendingService returns a pointer to a new UserPendingService
func NewUserPendingService(db *sql.DB) *UserPendingService {
	return &UserPendingService{
		userPendingRepo: repository.NewUserPendingRepository(db),
	}
}

// Insert for inserting forgotten user pending
func (s *UserPendingService) Insert(pending model.UserRegistrationPending) (int, error) {

	// Only hash if there is a password
	if pending.Password.Valid {
		// First hash the password for more security
		hashedPass, err := util.GenerateFromPassword(pending.Password.String)
		if err != nil {
			return -1, err
		}

		pending.Password = sql.NullString{
			String: hashedPass,
			Valid:  hashedPass != "",
		}
	}

	return s.userPendingRepo.Insert(pending)
}

// InsertNewEmail for inserting new email
func (s *UserPendingService) InsertNewEmail(pending model.UserRegistrationPending) (int, error) {
	return s.userPendingRepo.InsertNewEmail(pending)
}

// FetchAll fetches all rows in users_pending
func (s *UserPendingService) FetchAll() ([]*model.UserRegistrationPending, error) {
	return s.userPendingRepo.FetchAll()
}

// FetchPassword fetches one password to one user in users_pending
func (s *UserPendingService) FetchPassword(id int) (string, error) {
	return s.userPendingRepo.FetchPassword(id)
}

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

// NewUserPendingService func
func NewUserPendingService(db *sql.DB) *UserPendingService {
	return &UserPendingService{
		userPendingRepo: repository.NewUserPendingRepository(db),
	}
}

// Insert func
func (s *UserPendingService) Insert(pending model.UserPending) (int, error) {

	var err error

	// First hash the password for more security
	pending.Password, err = util.GenerateFromPassword(pending.Password)
	if err != nil {
		return -1, err
	}
	return s.userPendingRepo.Insert(pending)
}

// FetchAll fetches all rows in users_pending
func (s *UserPendingService) FetchAll() ([]*model.UserPending, error) {
	return s.userPendingRepo.FetchAll()
}

// FetchPassword fetches one password to one user in users_pending
func (s *UserPendingService) FetchPassword(id int) (string, error) {
	return s.userPendingRepo.FetchPassword(id)
}

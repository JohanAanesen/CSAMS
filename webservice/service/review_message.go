package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// UserService struct
type ReviewMessageService struct {
	reviewMessageRepo *repository.ReviewMessageRepository
}

// NewUserService returns a pointer to a new UserService
func NewReviewMessageService(db *sql.DB) *ReviewMessageService {
	return &ReviewMessageService{
		reviewMessageRepo: repository.NewReviewMessageRepository(db),
	}
}
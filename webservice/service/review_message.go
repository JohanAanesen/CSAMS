package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// ReviewMessageService struct
type ReviewMessageService struct {
	reviewMessageRepo *repository.ReviewMessageRepository
}

// NewReviewMessageService returns a pointer to a new UserService
func NewReviewMessageService(db *sql.DB) *ReviewMessageService {
	return &ReviewMessageService{
		reviewMessageRepo: repository.NewReviewMessageRepository(db),
	}
}

// FetchAllForAssignmentUser func
func (s *ReviewMessageService) FetchAllForAssignmentUser(assignmentID int, userID int) ([]*model.ReviewMessage, error){
	return s.reviewMessageRepo.FetchAllForAssignmentUser(assignmentID, userID)
}

// Insert func
func (s *ReviewMessageService) Insert(message model.ReviewMessage) error{
	return s.reviewMessageRepo.Insert(message)
}
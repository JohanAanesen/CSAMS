package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// PeerReviewService struct
type PeerReviewService struct {
	peerReviewRepo *repository.PeerReviewRepository
}

// NewPeerReviewService returns a pointer to a new PeerReviewService
func NewPeerReviewService(db *sql.DB) *PeerReviewService {
	return &PeerReviewService{
		peerReviewRepo: repository.NewPeerReviewRepository(db),
	}
}

// Insert a new peer review
func (s *PeerReviewService) Insert(assignmentID int, userID int, targetUserID int) (bool, error) {
	return s.peerReviewRepo.Insert(assignmentID, userID, targetUserID)
}

// TargetExists checks if the target exist in the table
func (s *PeerReviewService) TargetExists(assignmentID int, userID int) (bool, error) {
	return s.peerReviewRepo.TargetExists(assignmentID, userID)
}

// FetchAllFromAssignment from the database
func (s *PeerReviewService) FetchAllFromAssignment(assignmentID int) ([]*model.PeerReview, error) {
	return s.peerReviewRepo.FetchPeerReviewsFromAssignment(assignmentID)
}

// FetchReviewTargetsToUser from the database
func (s *PeerReviewService) FetchReviewTargetsToUser(userID int, assignmentID int) ([]*model.PeerReview, error) {
	return s.peerReviewRepo.FetchReviewTargetsToUser(userID, assignmentID)
}

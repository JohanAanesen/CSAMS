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

// NewPeerReviewService func
func NewPeerReviewService(db *sql.DB) *PeerReviewService {
	return &PeerReviewService{
		peerReviewRepo: repository.NewPeerReviewRepository(db),
	}
}

// Insert func
func (s *PeerReviewService) Insert(assignmentID int, userID int, targetUserID int) (bool, error) {
	return s.peerReviewRepo.Insert(assignmentID, userID, targetUserID)
}

// TargetExists checks if the target exist in the table
func (s *PeerReviewService) TargetExists(assignmentID int, userID int) (bool, error) {
	return s.peerReviewRepo.TargetExists(assignmentID, userID)
}

// FetchAllFromAssignment func
func (s *PeerReviewService) FetchAllFromAssignment(assignmentID int) ([]*model.PeerReview, error) {
	return s.peerReviewRepo.FetchPeerReviewsFromAssignment(assignmentID)
}

// FetchPeerReviewsToUser
func (s *PeerReviewService) FetchReviewTargetsToUser(userID int, assignmentID int) ([]*model.PeerReview, error) {
	return s.peerReviewRepo.FetchReviewTargetsToUser(userID, assignmentID)
}

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

// TargetExists checks if the target exist in the table
func (s *PeerReviewService) TargetExists(assignmentID int, userID int) (bool, error) {
	return s.peerReviewRepo.TargetExists(assignmentID, userID)
}

// FetchAllFromAssignment func
func (s *PeerReviewService) FetchAllFromAssignment(assignmentID int) ([]*model.PeerReview, error) {
	result := make([]*model.PeerReview, 0)

	peerReviewPtr, err := s.peerReviewRepo.FetchPeerReviewsFromAssignment(assignmentID)
	if err != nil {
		return result, err
	}

	for _, peerReview := range peerReviewPtr {
		result = append(result, peerReview)
	}

	return result, err
}
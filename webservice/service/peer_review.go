package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/repositroy"
)

// PeerReviewService struct
type PeerReviewService struct {
	peerReviewRepo *repositroy.PeerReviewRepository
}

// NewPeerReviewService func
func NewPeerReviewService(db *sql.DB) *PeerReviewService {
	return &PeerReviewService{
		peerReviewRepo: repositroy.NewPeerReviewRepository(db),
	}
}

// TargetExists checks if the target exist in the table
func (s *PeerReviewService) TargetExists(assignmentID int, userID int) (bool, error) {
	return s.peerReviewRepo.TargetExists(assignmentID, userID)
}

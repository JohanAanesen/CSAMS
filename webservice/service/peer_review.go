package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
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

func (s *PeerReviewService) TargetExists(assignmentID int, userID int) (bool, error) {
	return s.peerReviewRepo.TargetExists(assignmentID, userID)
}

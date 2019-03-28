package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
)

// ReviewService struct
type PeerReviewService struct {
	reviewRepo *repositroy.ReviewRepository
}

// NewReviewService func
func NewPeerReviewService(db *sql.DB) *PeerReviewService {
	return &PeerReviewService{
		reviewRepo: repositroy.NewReviewRepository(db),
	}
}

// FetchAllFromAssignment func
func (s *PeerReviewService) FetchAllFromAssignment(assignmentID int) ([]*model.PeerReview, error) {
	result := make([]*model.PeerReview, 0)

	peerReviewPtr, err := s.reviewRepo.FetchPeerReviewsFromAssignment(assignmentID)
	if err != nil {
		return result, err
	}

	for _, peerReview := range peerReviewPtr {
		result = append(result, peerReview)
	}

	return result, err
}

package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
)

// ReviewAnswerService struct
type ReviewAnswerService struct {
	reviewAnswerRepo *repositroy.ReviewAnswerRepository
}

// NewReviewAnswerService func
func NewReviewAnswerService(db *sql.DB) *ReviewAnswerService {
	return &ReviewAnswerService{
		reviewAnswerRepo: repositroy.NewReviewAnswerRepository(db),
	}
}

// FetchForAssignment func
func (s *ReviewAnswerService) FetchForAssignment(assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForAssignment(assignmentID)
}

// FetchForTarget func
func (s *ReviewAnswerService) FetchForTarget(target, assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForTarget(target, assignmentID)
}

// FetchForReviewer func
func (s *ReviewAnswerService) FetchForReviewer(reviewer, assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForReviewer(reviewer, assignmentID)
}

// FetchReviewUsers func
func (s *ReviewAnswerService) FetchReviewUsers(target, assignmentID int) ([]int, error) {
	users := make([]int, 0)

	answers, err := s.FetchForTarget(target, assignmentID)
	if err != nil {
		return users, err
	}


	for _, answer := range answers {
		if !util.Contains(users, answer.UserReviewer) {
			users = append(users, answer.UserReviewer)
		}
	}

	return users, err
}

// FetchForReviewerAndTarget func
func (s *ReviewAnswerService) FetchForReviewerAndTarget(reviewer, target, assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForReviewerAndTarget(reviewer, target, assignmentID)
}

// Insert func
func (s *ReviewAnswerService) Insert(answer model.ReviewAnswer) (int, error) {
	return s.reviewAnswerRepo.Insert(answer)
}

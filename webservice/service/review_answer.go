package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repositroy"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
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

// FetchForUser func
func (s *ReviewAnswerService) FetchForUser(userID, assignmentID int) ([][]*model.ReviewAnswer, error) {
	result := make([][]*model.ReviewAnswer, 0)

	reviewers, err := s.FetchReviewUsers(userID, assignmentID)
	if err != nil {
		return result, err
	}

	for _, k := range reviewers {
		review, err := s.FetchForReviewerAndTarget(k, userID, assignmentID)
		if err != nil {
			return result, err
		}

		result = append(result, review)
	}

	return result, err
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

// HasBeenReviewed func
func (s *ReviewAnswerService) HasBeenReviewed(target, reviewer, assignmentID int) (bool, error) {
	temp, err := s.reviewAnswerRepo.FetchForReviewerAndTarget(reviewer, target, assignmentID)
	if err != nil {
		return false, err
	}

	return len(temp) > 0, err
}

// CountReviewsDone func
func (s *ReviewAnswerService) CountReviewsDone(userID, assignmentID int) (int, error) {
	return s.reviewAnswerRepo.CountReviewsDone(userID, assignmentID)
}

// Insert func
func (s *ReviewAnswerService) Insert(answer model.ReviewAnswer) (int, error) {
	return s.reviewAnswerRepo.Insert(answer)
}

package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
	"time"
)

// SubmissionAnswerService struct
type SubmissionAnswerService struct {
	submissionAnswerRepo *repository.SubmissionAnswerRepository
	reviewAnswerRepo     *repository.ReviewAnswerRepository
}

// NewSubmissionAnswerService func
func NewSubmissionAnswerService(db *sql.DB) *SubmissionAnswerService {
	return &SubmissionAnswerService{
		submissionAnswerRepo: repository.NewSubmissionAnswerRepository(db),
		reviewAnswerRepo:     repository.NewReviewAnswerRepository(db),
	}
}

// Fetch func
func (s *SubmissionAnswerService) Fetch(id int) (*model.SubmissionAnswer, error) {
	return s.submissionAnswerRepo.Fetch(id)
}

// FetchAll func
func (s *SubmissionAnswerService) FetchAll() ([]*model.SubmissionAnswer, error) {
	return s.submissionAnswerRepo.FetchAll()
}

// FetchAllFromAssignment func
func (s *SubmissionAnswerService) FetchAllFromAssignment(assID int) ([]*model.SubmissionAnswer, error) {
	return s.submissionAnswerRepo.FetchAllForAssignment(assID)
}

// FetchUsersDelivered func
func (s *SubmissionAnswerService) FetchUsersDeliveredFromAssignment(assID int) ([]int, error){
	return s.submissionAnswerRepo.FetchUsersDeliveredFromAssignment(assID)
}

// CountForAssignment func
func (s *SubmissionAnswerService) CountForAssignment(assignmentID int) (int, error) {
	return s.submissionAnswerRepo.CountForAssignment(assignmentID)
}

// HasUserSubmitted func
func (s *SubmissionAnswerService) HasUserSubmitted(assignmentID, userID int) (bool, error) {
	answers, err := s.submissionAnswerRepo.FetchAll()
	if err != nil {
		return false, err
	}

	for _, item := range answers {
		if item.AssignmentID == assignmentID && item.UserID == userID {
			return true, err
		}
	}

	return false, err
}

// FetchUserAnswers func
func (s *SubmissionAnswerService) FetchUserAnswers(userID, assignmentID int) ([]*model.SubmissionAnswer, error) {
	return s.submissionAnswerRepo.FetchAllForUserAndAssignment(userID, assignmentID)
}

// Insert func
func (s *SubmissionAnswerService) Insert(answers []*model.SubmissionAnswer) error {
	for _, item := range answers {
		_, err := s.submissionAnswerRepo.Insert(*item)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update func
func (s *SubmissionAnswerService) Update(answers []*model.SubmissionAnswer) error {
	for _, item := range answers {
		err := s.submissionAnswerRepo.Update(*item)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteFromAssignment func
func (s *SubmissionAnswerService) DeleteFromAssignment(assignmentID int) error {
	return s.submissionAnswerRepo.DeleteFromAssignment(assignmentID)
}

// FetchSubmittedTime func
func (s *SubmissionAnswerService) FetchSubmittedTime(userID, assignmentID int) (time.Time, bool, error) {
	return s.submissionAnswerRepo.FetchSubmittedTime(userID, assignmentID)
}

// Delete func
func (s *SubmissionAnswerService) Delete(assignmentID, userID int) error {
	err := s.submissionAnswerRepo.Delete(assignmentID, userID)
	if err != nil {
		return err
	}

	return s.reviewAnswerRepo.DeleteTarget(assignmentID, userID)
}

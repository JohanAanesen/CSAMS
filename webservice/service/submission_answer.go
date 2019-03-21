package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
)

// SubmissionAnswerService struct
type SubmissionAnswerService struct {
	submissionAnswerRepo *repositroy.SubmissionAnswerRepository
}

// NewSubmissionAnswerService func
func NewSubmissionAnswerService(db *sql.DB) *SubmissionAnswerService {
	return &SubmissionAnswerService{
		submissionAnswerRepo: repositroy.NewSubmissionAnswerRepository(db),
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

// Upload func
func (s *SubmissionAnswerService) Upload(answers []*model.SubmissionAnswer) error {
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
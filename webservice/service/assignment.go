package service

import (
	"database/sql"

	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// AssignmentService struct
type AssignmentService struct {
	assignmentRepo *repository.AssignmentRepository
}

// NewAssignmentService func
func NewAssignmentService(db *sql.DB) *AssignmentService {
	return &AssignmentService{
		assignmentRepo: repository.NewAssignmentRepository(db),
	}
}

// Fetch func
func (s *AssignmentService) Fetch(id int) (*model.Assignment, error) {
	return s.assignmentRepo.Fetch(id)
}

// FetchAll func
func (s *AssignmentService) FetchAll() ([]*model.Assignment, error) {
	return s.assignmentRepo.FetchAll()
}

// FetchFromCourse func
func (s *AssignmentService) FetchFromCourse(courseID int) ([]*model.Assignment, error) {
	result := make([]*model.Assignment, 0)

	assignments, err := s.assignmentRepo.FetchAll()
	if err != nil {
		return result, err
	}

	for _, item := range assignments {
		if item.CourseID == courseID {
			result = append(result, item)
		}
	}

	return result, err
}

// HasReview func
func (s *AssignmentService) HasReview(assignmentID int) (bool, error) {
	assignment, err := s.assignmentRepo.Fetch(assignmentID)
	if err != nil {
		return false, err
	}

	return assignment.ReviewID.Valid, err
}

// HasAutoValidation func
func (s *AssignmentService) HasAutoValidation(assignmentID int) (bool, error) {
	assignment, err := s.assignmentRepo.Fetch(assignmentID)
	if err != nil {
		return false, err
	}

	return assignment.ValidationID.Valid, err
}

// Insert func
func (s *AssignmentService) Insert(assignment model.Assignment) (int, error) {
	return s.assignmentRepo.Insert(assignment)
}

// Update func
func (s *AssignmentService) Update(assignment model.Assignment) error {
	return s.assignmentRepo.Update(assignment)
}

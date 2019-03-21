package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
)

// AssignmentService struct
type AssignmentService struct {
	assignmentRepo *repositroy.AssignmentRepository
}

// NewAssignmentService func
func NewAssignmentService(db *sql.DB) *AssignmentService {
	return &AssignmentService{
		assignmentRepo: repositroy.NewAssignmentRepository(db),
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

// Insert func
func (s *AssignmentService) Insert(assignment model.Assignment) (int, error) {
	return s.assignmentRepo.Insert(assignment)
}

// Update func
func (s *AssignmentService) Update(assignment model.Assignment) error {
	return s.assignmentRepo.Update(assignment)
}
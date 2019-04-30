package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// SubmissionService struct
type SubmissionService struct {
	submissionRepo *repository.SubmissionRepository
	formRepo       *repository.FormRepository
	fieldRepo      *repository.FieldRepository
	assignmentRepo *repository.AssignmentRepository
}

// NewSubmissionService func
func NewSubmissionService(db *sql.DB) *SubmissionService {
	return &SubmissionService{
		submissionRepo: repository.NewSubmissionRepository(db),
		formRepo:       repository.NewFormRepository(db),
		fieldRepo:      repository.NewFieldRepository(db),
		assignmentRepo: repository.NewAssignmentRepository(db),
	}
}

// FetchAll func
func (s *SubmissionService) FetchAll() ([]model.Submission, error) {
	result := make([]model.Submission, 0)

	reviewPtr, err := s.submissionRepo.FetchAll()
	if err != nil {
		return result, err
	}

	formsPtr, err := s.formRepo.FetchAll()
	if err != nil {
		return result, err
	}

	for _, submission := range reviewPtr {
		for _, form := range formsPtr {
			if submission.FormID == form.ID {
				submission.Form = *form
			}
		}

		result = append(result, *submission)
	}

	return result, err
}

// Fetch func
func (s *SubmissionService) Fetch(id int) (*model.Submission, error) {
	return s.submissionRepo.Fetch(id)
}

// FetchFromFormID fetches submission from form id
func (s *SubmissionService) FetchFromFormID(formID int) (*model.Submission, error) {
	return s.submissionRepo.FetchFromFormID(formID)
}

// FetchFromAssignment func
func (s *SubmissionService) FetchFromAssignment(assignmentID int) (*model.Submission, error) {
	result := model.Submission{}

	assignment, err := s.assignmentRepo.Fetch(assignmentID)
	if err != nil {
		return &result, err
	}

	temp, err := s.submissionRepo.Fetch(int(assignment.SubmissionID.Int64))
	if err != nil {
		return &result, err
	}

	form, err := s.formRepo.Fetch(temp.FormID)
	if err != nil {
		return &result, err
	}

	fields, err := s.fieldRepo.FetchAllFromForm(form.ID)
	if err != nil {
		return &result, err
	}

	for _, field := range fields {
		form.Fields = append(form.Fields, *field)
	}

	temp.Form = *form

	return temp, err
}

// Insert func
func (s *SubmissionService) Insert(form model.Form) (int, error) {
	return s.submissionRepo.Insert(form)
}

// Update func
func (s *SubmissionService) Update(form model.Form) error {
	err := s.formRepo.Update(form.ID, &form)
	if err != nil {
		return err
	}

	err = s.fieldRepo.DeleteAll(form.ID)
	if err != nil {
		return err
	}

	for _, field := range form.Fields {
		field.FormID = form.ID

		_, err = s.fieldRepo.Insert(&field)
		if err != nil {
			return err
		}
	}

	return err
}

// Delete func
func (s *SubmissionService) Delete(id int) error {
	err := s.submissionRepo.DeleteWithFormID(id)
	if err != nil {
		return err
	}

	err = s.fieldRepo.DeleteAll(id)
	if err != nil {
		return err
	}

	err = s.formRepo.Delete(id)
	if err != nil {
		return err
	}

	return err
}

// IsUsed func
func (s *SubmissionService) IsUsed(formID int) (bool, error) {
	return s.submissionRepo.IsUsed(formID)
}

// UsedInAssignment func
func (s *SubmissionService) UsedInAssignment(formID int) (int, error) {
	return s.submissionRepo.UsedInAssignment(formID)
}

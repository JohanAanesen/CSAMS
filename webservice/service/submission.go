package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
)

// SubmissionService struct
type SubmissionService struct {
	submissionRepo *repositroy.SubmissionRepository
	formRepo       *repositroy.FormRepository
	fieldRepo      *repositroy.FieldRepository
}

// NewSubmissionService func
func NewSubmissionService(db *sql.DB) *SubmissionService {
	return &SubmissionService{
		submissionRepo: repositroy.NewSubmissionRepository(db),
		formRepo:       repositroy.NewFormRepository(db),
		fieldRepo:      repositroy.NewFieldRepository(db),
	}
}

// FetchAll func
func (rs *SubmissionService) FetchAll() ([]model.Submission, error) {
	result := make([]model.Submission, 0)

	reviewPtr, err := rs.submissionRepo.FetchAll()
	if err != nil {
		return result, err
	}

	formsPtr, err := rs.formRepo.FetchAll()
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

// Insert func
func (rs *SubmissionService) Insert(form model.Form) (int, error) {
	return rs.submissionRepo.Insert(form)
}

// Update func
func (rs *SubmissionService) Update(form model.Form) error {
	err := rs.formRepo.Update(form.ID, &form)
	if err != nil {
		return err
	}

	err = rs.fieldRepo.DeleteAll(form.ID)
	if err != nil {
		return err
	}

	for _, field := range form.Fields {
		field.FormID = form.ID

		_, err = rs.fieldRepo.Insert(&field)
		if err != nil {
			return err
		}
	}

	return err
}

// Delete func
func (rs *SubmissionService) Delete(id int) error {
	err := rs.submissionRepo.Delete(id)
	if err != nil {
		return err
	}

	err = rs.fieldRepo.DeleteAll(id)
	if err != nil {
		return err
	}

	err = rs.formRepo.Delete(id)
	if err != nil {
		return err
	}

	return err
}

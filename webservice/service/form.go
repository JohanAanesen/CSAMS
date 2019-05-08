package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// FormService struct
type FormService struct {
	formRepo  *repository.FormRepository
	fieldRepo *repository.FieldRepository
}

// NewFormService returns a pointer to a new FormService
func NewFormService(db *sql.DB) *FormService {
	return &FormService{
		formRepo:  repository.NewFormRepository(db),
		fieldRepo: repository.NewFieldRepository(db),
	}
}

// FetchAll forms from the database
func (fs *FormService) FetchAll() ([]*model.Form, error) {
	forms, err := fs.formRepo.FetchAll()
	return forms, err
}

// Fetch a single forms from the database
func (fs *FormService) Fetch(id int) (*model.Form, error) {
	form, err := fs.formRepo.Fetch(id)
	if err != nil {
		return nil, err
	}
	fields, err := fs.fieldRepo.FetchAllFromForm(id)
	if err != nil {
		return nil, err
	}

	for _, field := range fields {
		form.Fields = append(form.Fields, *field)
	}

	return form, err
}

// Insert a form to the database
func (fs *FormService) Insert(form model.Form) (int, error) {
	return fs.formRepo.Insert(&form)
}

// Update a form in the database
func (fs *FormService) Update(id int, form model.Form) error {
	return fs.formRepo.Update(id, &form)
}

// Delete a form from the database based on the ID
func (fs *FormService) Delete(id int) error {
	return fs.formRepo.Delete(id)
}

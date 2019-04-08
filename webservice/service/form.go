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

// NewFormService func
func NewFormService(db *sql.DB) *FormService {
	return &FormService{
		formRepo:  repository.NewFormRepository(db),
		fieldRepo: repository.NewFieldRepository(db),
	}
}

// FetchAll func
func (fs *FormService) FetchAll() ([]*model.Form, error) {
	forms, err := fs.formRepo.FetchAll()
	return forms, err
}

// Fetch func
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

// Insert func
func (fs *FormService) Insert(form model.Form) (int, error) {
	return fs.formRepo.Insert(&form)
}

// Update func
func (fs *FormService) Update(id int, form model.Form) error {
	return fs.formRepo.Update(id, &form)
}

// Delete func
func (fs *FormService) Delete(id int) error {
	return fs.formRepo.Delete(id)
}

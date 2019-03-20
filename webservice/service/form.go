package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
)

// FormService struct
type FormService struct {
	formRepo  *repositroy.FormRepository
	fieldRepo *repositroy.FieldRepository
}

// NewFormService func
func NewFormService(db *sql.DB) *FormService {
	return &FormService{
		formRepo:  repositroy.NewFormRepository(db),
		fieldRepo: repositroy.NewFieldRepository(db),
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
	fields, err := fs.fieldRepo.FetchAll()
	if err != nil {
		return nil, err
	}

	for _, field := range fields {
		if field.FormID == form.ID {
			form.Fields = append(form.Fields, *field)
		}
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

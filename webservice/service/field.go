package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// FieldService struct
type FieldService struct {
	repo *repository.FieldRepository
}

// NewFieldService returns a new pointer to a FieldService
func NewFieldService(db *sql.DB) *FieldService {
	return &FieldService{
		repo: repository.NewFieldRepository(db),
	}
}

// FetchAllWithFormID fetches all fields from a form with a given ID
func (fs *FieldService) FetchAllWithFormID(formID int) ([]model.Field, error) {
	ptr, err := fs.repo.FetchAll()

	fields := make([]model.Field, 0)

	for _, field := range ptr {
		if field.FormID == formID {
			fields = append(fields, *field)
		}
	}

	return fields, err
}

// FetchAll all fields
func (fs *FieldService) FetchAll() ([]model.Field, error) {
	ptr, err := fs.repo.FetchAll()

	fields := make([]model.Field, 0)
	for _, field := range ptr {
		fields = append(fields, *field)
	}

	return fields, err
}

// Fetch a sinle field
func (fs *FieldService) Fetch(id int) (*model.Field, error) {
	field, err := fs.repo.Fetch(id)

	return field, err
}

// Insert field to database
func (fs *FieldService) Insert(field model.Field) (int, error) {
	return fs.repo.Insert(&field)
}

// Update field in the database
func (fs *FieldService) Update(id int, field model.Field) error {
	return fs.repo.Update(id, &field)
}

// Delete field from the database
func (fs *FieldService) Delete(id int) error {
	return fs.repo.DeleteAll(id)
}

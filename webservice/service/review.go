package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
	_ "github.com/go-sql-driver/mysql" //database driver
)

// ReviewService struct
type ReviewService struct {
	reviewRepo *repositroy.ReviewRepository
	formRepo   *repositroy.FormRepository
	fieldRepo  *repositroy.FieldRepository
}

// NewReviewService func
func NewReviewService(db *sql.DB) *ReviewService {
	return &ReviewService{
		reviewRepo: repositroy.NewReviewRepository(db),
		formRepo:   repositroy.NewFormRepository(db),
		fieldRepo:  repositroy.NewFieldRepository(db),
	}
}

// FetchAll func
func (rs *ReviewService) FetchAll() ([]model.Review, error) {
	result := make([]model.Review, 0)

	reviewPtr, err := rs.reviewRepo.FetchAll()
	if err != nil {
		return result, err
	}

	formsPtr, err := rs.formRepo.FetchAll()
	if err != nil {
		return result, err
	}

	for _, review := range reviewPtr {
		for _, form := range formsPtr {
			if review.FormID == form.ID {
				review.Form = *form
			}
		}

		result = append(result, *review)
	}

	return result, err
}

// Insert func
func (rs *ReviewService) Insert(form model.Form) (int, error) {
	return rs.reviewRepo.Insert(form)
}

// Update func
func (rs *ReviewService) Update(form model.Form) error {
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
func (rs *ReviewService) Delete(id int) error {
	err := rs.reviewRepo.Delete(id)
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

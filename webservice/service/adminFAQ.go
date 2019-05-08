package service

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// FAQService struct
type FAQService struct {
	faqRepo *repository.FAQRepository
}

// NewFAQService func
func NewFAQService(db *sql.DB) *FAQService {
	return &FAQService{
		faqRepo: repository.NewFAQRepository(db),
	}
}

// InsertNew inserts a new faq in db, if there isn't one in db already
func (s *FAQService) InsertNew() error {

	faq, err := s.faqRepo.Fetch()
	if err != nil {
		return err
	}

	// Check if there isn't a faq in db already
	if faq.Questions != "" {
		return errors.New("error: there already exist an faq in db")
	}

	// Faq doesn't exist in db, create new
	return s.faqRepo.InsertNew()
}

// Fetch fetches the one faq in db
func (s *FAQService) Fetch() (*model.Faq, error) {
	return s.faqRepo.Fetch()
}

// Update updates the one faq in db
func (s *FAQService) Update(questions string) error {
	return s.faqRepo.Update(questions)
}

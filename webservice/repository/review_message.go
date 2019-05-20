package repository

import "database/sql"

// UserRepository struct
type ReviewMessageRepository struct {
	db *sql.DB
}

// NewUserRepository return a pointer to a new UserRepository
func NewReviewMessageRepository(db *sql.DB) *ReviewMessageRepository {
	return &ReviewMessageRepository{
		db: db,
	}
}
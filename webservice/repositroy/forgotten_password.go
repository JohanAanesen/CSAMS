package repositroy

import "database/sql"

// ForgottenPassRepository struct
type ForgottenPassRepository struct {
	db *sql.DB
}

// NewForgottenPassRepository func
func NewForgottenPassRepository(db *sql.DB) *ForgottenPassRepository {
	return &ForgottenPassRepository{
		db: db,
	}
}

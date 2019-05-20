package repository

import "database/sql"

// UserRepository struct
type NotificationRepository struct {
	db *sql.DB
}

// NewUserRepository return a pointer to a new UserRepository
func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

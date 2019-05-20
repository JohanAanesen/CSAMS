package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// UserService struct
type NotificationService struct {
	notificationRepo *repository.NotificationRepository
}

// NewUserService returns a pointer to a new UserService
func NewNotificationService(db *sql.DB) *NotificationService {
	return &NotificationService{
		notificationRepo: repository.NewNotificationRepository(db),
	}
}

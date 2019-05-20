package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// UserService struct
type NotificationService struct {
	notificationRepo *repository.NotificationRepository
}

// NewUserService returns a pointer to a new NotificationService
func NewNotificationService(db *sql.DB) *NotificationService {
	return &NotificationService{
		notificationRepo: repository.NewNotificationRepository(db),
	}
}

func (s *NotificationService) FetchAllForUser(UserID int) ([]*model.Notification, error) {
	return s.notificationRepo.FetchAllForUser(UserID)
}
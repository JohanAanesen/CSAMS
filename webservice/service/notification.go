package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// NotificationService struct
type NotificationService struct {
	notificationRepo *repository.NotificationRepository
}

// NewNotificationService returns a pointer to a new NotificationService
func NewNotificationService(db *sql.DB) *NotificationService {
	return &NotificationService{
		notificationRepo: repository.NewNotificationRepository(db),
	}
}

// FetchAllForUser func
func (s *NotificationService) FetchAllForUser(UserID int) ([]*model.Notification, error) {
	return s.notificationRepo.FetchAllForUser(UserID)
}

// FetchNotificationForUser func
func (s *NotificationService) FetchNotificationForUser(UserID int, NotificationID int) (*model.Notification, error) {
	return s.notificationRepo.FetchNotificationForUser(UserID, NotificationID)
}

// CountUnreadNotifications func
func (s *NotificationService) CountUnreadNotifications(UserID int) (int, error) {
	return s.notificationRepo.CountUnreadNotifications(UserID)
}

// MarkAsRead func
func (s *NotificationService) MarkAsRead(NotificationID int) error {
	return s.notificationRepo.MarkAsRead(NotificationID)
}

// Insert func
func (s *NotificationService) Insert(notification model.Notification) (int, error) {
	return s.notificationRepo.Insert(notification)
}

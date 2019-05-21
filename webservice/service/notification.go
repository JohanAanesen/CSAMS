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

// FetchAllForUser func
func (s *NotificationService) FetchAllForUser(UserID int) ([]*model.Notification, error) {
	return s.notificationRepo.FetchAllForUser(UserID)
}

// FetchNotificationForUSer func
func (s *NotificationService) FetchNotificationForUSer(UserID int, NotificationID int) (*model.Notification, error){
	return s.notificationRepo.FetchNotificationForUSer(UserID, NotificationID)
}

// CountUnreadNotifications func
func (s *NotificationService) CountUnreadNotifications(UserID int) (int, error){
	return s.notificationRepo.CountUnreadNotifications(UserID)
}

// MarkAsRead func
func (s *NotificationService) MarkAsRead(NotificationID int) error{
	return s.notificationRepo.MarkAsRead(NotificationID)
}
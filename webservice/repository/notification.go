package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
)

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

// FetchAllForUser func
func (repo *NotificationRepository) FetchAllForUser(userID int) ([]*model.Notification, error) {
	result := make([]*model.Notification, 0)

	query := "SELECT id, user_id, url, message, active FROM notifications WHERE user_id = ?"

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.Notification{}

		err = rows.Scan(&temp.ID, &temp.UserID, &temp.URL,
			&temp.Message, &temp.Active)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchNotificationForUser function fetches single notification from userid and notifyID
func (repo *NotificationRepository) FetchNotificationForUSer(userID int, notificationID int) (*model.Notification, error) {
	result := model.Notification{}

	query := "SELECT id, user_id, url, message, active FROM notifications WHERE user_id = ? AND id = ?"

	rows, err := repo.db.Query(query, userID, notificationID)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.UserID, &result.URL,
			&result.Message, &result.Active)
		if err != nil {
			return &result, err
		}
	}

	return &result, err
}

// Insert func
func (repo *NotificationRepository) Insert(notification model.Notification) (int, error) {
	var id int64

	query := "INSERT INTO notifications (user_id, url, message) VALUES (?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	rows, err := tx.Exec(query, notification.UserID, notification.URL, notification.Message)
	if err != nil {
		_ = tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return int(id), err
}

func (repo *NotificationRepository) MarkAsRead(notificationID int) error {
	query := "UPDATE notifications SET active = 0 WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, notificationID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (repo *NotificationRepository) CountUnreadNotifications(userID int) (int, error) {
	count := 0

	query := "SELECT SUM(active) FROM notifications WHERE user_id = ?"

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return count, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, err
}

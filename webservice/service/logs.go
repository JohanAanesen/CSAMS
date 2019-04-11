package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// LogsService struct
type LogsService struct {
	logsRepo *repository.LogsRepository
}

// NewLogsService func
func NewLogsService(db *sql.DB) *LogsService {
	return &LogsService{
		logsRepo: repository.NewLogsRepository(db),
	}
}

// FetchAll fetches all logs
func (s *LogsService) FetchAll() ([]*model.Logs, error) {
	return s.logsRepo.FetchAll()
}

// InsertNewUser inserts a new user log
func (s *LogsService) InsertNewUser(userID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.NewUser,
	}

	return s.logsRepo.Insert(logx)
}

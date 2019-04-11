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

// FetchAll func
func (s *LogsService) FetchAll() ([]*model.Logs, error) {
	return s.logsRepo.FetchAll()
}

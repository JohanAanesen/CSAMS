package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
)

// ForgottenPassService struct
type ForgottenPassService struct {
	assignmentRepo *repositroy.ForgottenPassRepository
}

// NewForgottenPassService func
func NewForgottenPassService(db *sql.DB) *ForgottenPassService {
	return &ForgottenPassService{
		assignmentRepo: repositroy.NewForgottenPassRepository(db),
	}
}

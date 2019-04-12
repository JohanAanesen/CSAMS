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

// InsertChangeEmail inserts a change email log
func (s *LogsService) InsertChangeEmail(userID int, oldValue string, newValue string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.ChangeEmail,
	}

	// Add oldValue to struct
	logx.OldValue = sql.NullString{
		String: oldValue,
		Valid:  oldValue != "",
	}

	// Add newValue to struct
	logx.NewValue = sql.NullString{
		String: newValue,
		Valid:  newValue != "",
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangeFAQ inserts a change FAQ log
func (s *LogsService) InsertChangeFAQ(userID int, oldValue string, newValue string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.ChangeAdminFAQ,
	}

	// Add oldValue to struct
	logx.OldValue = sql.NullString{
		String: oldValue,
		Valid:  oldValue != "",
	}

	// Add newValue to struct
	logx.NewValue = sql.NullString{
		String: newValue,
		Valid:  newValue != "",
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangePassword inserts a change password log
func (s *LogsService) InsertChangePassword(userID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.ChangePassword,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAssignment inserts a change password log
func (s *LogsService) InsertAssignment(userID int, assignmentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.DeliveredAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentId = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangeAssignment inserts a change password log
func (s *LogsService) InsertChangeAssignment(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.UpdateAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentId = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertDeleteAssignment inserts a change password log
func (s *LogsService) InsertDeleteAssignment(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.DeleteAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentId = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAllReviewsDone is for when all reviews to one submission is done
func (s *LogsService) InsertAllReviewsDone(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.PeerReviewDone,
	}

	// Add assignmentID to struct
	logx.AssignmentId = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertFinishedOnePeerReview is for when one user has finished peer reviewing another users submission
func (s *LogsService) InsertFinishedOnePeerReview(userID int, assignmentID int, submissionID int, affectedUserID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.FinishedOnePeerReview,
	}

	// Add assignmentID to struct
	logx.AssignmentId = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	// Add affectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(affectedUserID),
		Valid: affectedUserID != 0,
	}
	return s.logsRepo.Insert(logx)
}

// InsertCourse inserts a new course log
func (s *LogsService) InsertCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.CreatedCourse,
	}

	// Add assignmentID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertJoinCourse inserts a new join course log
func (s *LogsService) InsertJoinCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.JoinedCourse,
	}

	// Add assignmentID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

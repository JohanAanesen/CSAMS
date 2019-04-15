package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
	"strings"
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

// InsertUpdateFAQ inserts a updated FAQ log
func (s *LogsService) InsertUpdateFAQ(userID int, oldValue string, newValue string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateFAQ,
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

// InsertChangePasswordEmail inserts a change password with email log
func (s *LogsService) InsertChangePasswordEmail(userID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.ChangePasswordEmail,
	}

	return s.logsRepo.Insert(logx)
}

// InsertSubmission inserts a new user submission log
func (s *LogsService) InsertSubmission(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.CreateSubmission,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
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

// InsertUpdateSubmission inserts a update user submission log
func (s *LogsService) InsertUpdateSubmission(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.UpdateSubmission,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
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

// InsertDeleteSubmission inserts a delete user submission log
func (s *LogsService) InsertDeleteSubmission(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.DeleteSubmission,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
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

// InsertAdminCreateAssignment inserts a create assignment log
func (s *LogsService) InsertAdminCreateAssignment(userID int, assignmentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreateAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminUpdateAssignment inserts a change password log
func (s *LogsService) InsertAdminUpdateAssignment(userID int, assignmentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminDeleteAssignment inserts a change password log
func (s *LogsService) InsertAdminDeleteAssignment(userID int, assignmentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertFinishedOnePeerReview is for when one user has finished peer reviewing another users submission
func (s *LogsService) InsertFinishedOnePeerReview(userID int, assignmentID int, affectedUserID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.FinishedOnePeerReview,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add affectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(affectedUserID),
		Valid: affectedUserID != 0,
	}
	return s.logsRepo.Insert(logx)
}

// InsertUpdateOnePeerReview is for when one user has updated peer review
func (s *LogsService) InsertUpdateOnePeerReview(userID int, assignmentID int, affectedUserID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.UpdateOnePeerReview,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
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
		Activity: model.AdminCreatedCourse,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertUpdateCourse inserts a new update course log
func (s *LogsService) InsertUpdateCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateCourse,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertDeleteCourse inserts a new delete course log
func (s *LogsService) InsertDeleteCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteCourse,
	}

	// Add courseID to struct
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

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertLeftCourse inserts a new join course log
func (s *LogsService) InsertLeftCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.LeftCourse,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminSubmissionForm inserts a new submission form
func (s *LogsService) InsertAdminSubmissionForm(userID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreateSubmissionForm,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminUpdateSubmissionForm inserts a new submission form
func (s *LogsService) InsertAdminUpdateSubmissionForm(userID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateSubmissionForm,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminDeleteSubmissionForm inserts a new submission form
func (s *LogsService) InsertAdminDeleteSubmissionForm(userID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteSubmissionForm,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminReviewForm inserts a new review form
func (s *LogsService) InsertAdminReviewForm(userID int, reviewID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreateReviewForm,
	}

	// Add reviewID to struct
	logx.ReviewID = sql.NullInt64{
		Int64: int64(reviewID),
		Valid: reviewID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminUpdateReviewForm inserts a new review form
func (s *LogsService) InsertAdminUpdateReviewForm(userID int, reviewID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateReviewForm,
	}

	// Add reviewID to struct
	logx.ReviewID = sql.NullInt64{
		Int64: int64(reviewID),
		Valid: reviewID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminDeleteReviewForm inserts a new review form
func (s *LogsService) InsertAdminDeleteReviewForm(userID int, reviewID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteReviewForm,
	}

	// Add reviewID to struct
	logx.ReviewID = sql.NullInt64{
		Int64: int64(reviewID),
		Valid: reviewID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertEmailStudents inserts a new email students log
func (s *LogsService) InsertEmailStudents(userID int, courseID int, emails []string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminEmailCourseStudents,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	emailsString := strings.Join(emails, ",")

	// Add newvalue to struct, all the emails sent to
	logx.NewValue = sql.NullString{
		String: emailsString,
		Valid:  emailsString != "",
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminRemoveUserFromCourse inserts a remove student from course log
func (s *LogsService) InsertAdminRemoveUserFromCourse(userID int, courseID int, studentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminRemoveUserFromCourse,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	// Add AffectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(studentID),
		Valid: studentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminChangeUserPassword inserts a change student password log
func (s *LogsService) InsertAdminChangeUserPassword(userID int, studentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminChangeStudentPassword,
	}

	// Add AffectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(studentID),
		Valid: studentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminCreateSubmissionForUser inserts a create submission for student log
func (s *LogsService) InsertAdminCreateSubmissionForUser(userID int, assignmentID int, submissionID int, studentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreateSubmissionForUser,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	// Add AffectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(studentID),
		Valid: studentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminUpdateSubmissionForUser inserts a update submission for student log
func (s *LogsService) InsertAdminUpdateSubmissionForUser(userID int, assignmentID int, submissionID int, studentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateSubmissionForUser,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	// Add AffectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(studentID),
		Valid: studentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminDeleteSubmissionForUser inserts a delete submission for student log
func (s *LogsService) InsertAdminDeleteSubmissionForUser(userID int, assignmentID int, submissionID int, studentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteSubmissionForUser,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	// Add AffectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(studentID),
		Valid: studentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

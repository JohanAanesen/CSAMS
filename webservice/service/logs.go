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

// InsertAdminCourse inserts a new course log
func (s *LogsService) InsertAdminCourse(userID int, courseID int) error {

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

// InsertAdminUpdateCourse inserts a new update course log
func (s *LogsService) InsertAdminUpdateCourse(userID int, courseID int) error {

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

// InsertAdminDeleteCourse inserts a new delete course log
func (s *LogsService) InsertAdminDeleteCourse(userID int, courseID int) error {

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

// InsertAdminCreateGroup inserts a create group log
func (s *LogsService) InsertAdminCreateGroup(userID int, groupID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreateGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminDeleteGroup inserts a delete group log
func (s *LogsService) InsertAdminDeleteGroup(userID int, groupID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminEditGroupName inserts a edit group name log
func (s *LogsService) InsertAdminEditGroupName(userID int, groupID int, oldValue string, newValue string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminEditGroupName,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
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

// InsertAdminAddUserToGroup inserts a add user to group log
func (s *LogsService) InsertAdminAddUserToGroup(userID int, groupID int, affectedUserID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminAddUserToGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	// Add affectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(affectedUserID),
		Valid: affectedUserID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAdminRemoveUserFromGroup inserts a remove user from group log
func (s *LogsService) InsertAdminRemoveUserFromGroup(userID int, groupID int, affectedUserID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminRemoveUserFromGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	// Add affectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(affectedUserID),
		Valid: affectedUserID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertCreateGroup inserts a create group log
func (s *LogsService) InsertCreateGroup(userID int, groupID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.CreateGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertDeleteGroup inserts a delete group log
func (s *LogsService) InsertDeleteGroup(userID int, groupID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.DeleteGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertJoinGroup inserts a join group log
func (s *LogsService) InsertJoinGroup(userID int, groupID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.JoinGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertLeftGroup inserts a left group log
func (s *LogsService) InsertLeftGroup(userID int, groupID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.LeftGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertEditGroupName inserts a edit group name log
func (s *LogsService) InsertEditGroupName(userID int, groupID int, oldValue string, newValue string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.EditGroupName,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
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

// InsertRemoveUserFromGroup inserts a removed user from course log
func (s *LogsService) InsertRemoveUserFromGroup(userID int, groupID int, removedUser int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.KickedFromGroup,
	}

	// Add groupID to struct
	logx.GroupID = sql.NullInt64{
		Int64: int64(groupID),
		Valid: groupID != 0,
	}

	// Add affectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(removedUser),
		Valid: removedUser != 0,
	}

	return s.logsRepo.Insert(logx)
}

package service

import "database/sql"

// Services struct collect all services into one struct, for less code in the handlers
type Services struct {
	Assignment       *AssignmentService
	Course           *CourseService
	Field            *FieldService
	Form             *FormService
	Review           *ReviewService
	ReviewAnswer     *ReviewAnswerService
	PeerReview       *PeerReviewService
	Submission       *SubmissionService
	SubmissionAnswer *SubmissionAnswerService
	User             *UserService
	Validation       *ValidationService
	UserPending      *UserPendingService
	GroupService     *GroupService
	Logs             *LogsService
	FAQ              *FAQService
	Notification     *NotificationService
	ReviewMessage    *ReviewMessageService
}

// NewServices returns a pointer to a new Services
func NewServices(db *sql.DB) *Services {
	return &Services{
		Assignment:       NewAssignmentService(db),
		Course:           NewCourseService(db),
		Field:            NewFieldService(db),
		Form:             NewFormService(db),
		Review:           NewReviewService(db),
		ReviewAnswer:     NewReviewAnswerService(db),
		PeerReview:       NewPeerReviewService(db),
		Submission:       NewSubmissionService(db),
		SubmissionAnswer: NewSubmissionAnswerService(db),
		User:             NewUserService(db),
		Validation:       NewValidationService(db),
		UserPending:      NewUserPendingService(db),
		GroupService:     NewGroupService(db),
		Logs:             NewLogsService(db),
		FAQ:              NewFAQService(db),
		Notification:     NewNotificationService(db),
		ReviewMessage:    NewReviewMessageService(db),
	}
}

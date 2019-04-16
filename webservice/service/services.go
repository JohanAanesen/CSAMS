package service

import "database/sql"

// Services struct
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
	Logs             *LogsService
	FAQ              *FAQService
}

// NewServices func
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
		Logs:             NewLogsService(db),
		FAQ:              NewFAQService(db),
	}
}

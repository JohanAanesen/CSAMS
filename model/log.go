package model

// activity enum for keeping track of log activity
type activity string

// Enum for logs
const (
	ChangeName          activity = "CHANGE-NAME"                            // User changed name
	ChangeEmail         activity = "CHANGE-EMAIL"                           // User changed email
	ChangePassword      activity = "CHANGE-PASSWORD"                        // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	DeliveredAssignment activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	FinishedPeerReview  activity = "FINISHED-PEER-REVIEWING"                // User is done peer reviewing two assignments
	PeerReviewDone      activity = "PEER-REVIEW-IS-DONE-FOR-ONE-ASSIGNMENT" // Users assignment is finished peer-reviewd
	JoinedCourse        activity = "STUDENT-JOINED-COURSE"                  // User joined course
	CreatedCourse       activity = "COURSE-CREATED"                         // Course is created
	CreatAssignment     activity = "ASSIGNMENT-CREATED"                     // Assignment is created
)

// Log struct to hold log-data
type Log struct {
	UserID       int      // [NOT NULL][all] User identification
	Activity     activity // [NOT NULL][all] User activity
	IsTeacher    bool     // [NULLABLE][later user] Says if the user is student or teacher (This is later checked from database)
	AssignmentID int      // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone/CreatAssignment] ID to relative assignment
	CourseID     int      // [NULLABLE][JoinedCourse/CreatedCourse] ID to relative course
	SubmissionID int      // [NULLABLE][DeliveredAssignment/FinishedPeerReview/PeerReviewDone] ID to relative submission
	OldValue     string   // [NULLABLE][ChangeName/ChangeEmail] Value before changing name/email
	NewValue     string   // [NULLABLE][ChangeName/ChangeEmail] Value after changing name/email
}

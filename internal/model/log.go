package model

// activity enum for keeping track of log activity
type activity string
const (
	ChangeName          activity = "CHANGE-NAME"                            // User changed name
	ChangeEmail         activity = "CHANGE-EMAIL"                           // User changed email
	ChangePassword      activity = "CHANGE-PASSWORD"                        // User changed password (DO NOT SHOW OLD/NEW PASSWORD IN LOG)
	DeliveredAssignment activity = "ASSIGNMENT-DELIVERED"                   // User delivered assignment
	FinishedPeerReview  activity = "FINISHED-PEER-REVIEWING"                // User is done peer reviewing two assignments
	PeerReviewDone      activity = "PEER-REVIEW-IS-DONE-FOR-ONE-ASSIGNMENT" // Users assignment is finished peer-reviewd
	JoinedCousrse       activity = "STUDENT-JOINED-COURSE"                  // User joined course
	CreatedCourse       activity = "COURSE-CREATED"                         // Course is created
	CreatAssignment     activity = "ASSIGNMENT-CREATED"                     // Assignment is created
)

// Log struct to hold log-data
type Log struct {
	UserID           int      // [NOT NULL] User identification
	Activity         activity // [NOT NULL] User activity
	AssignmentID     int      // [NULLABLE] ID to relative assignment
	CourseID         int      // [NULLABLE] ID to relative course
	UserAssignmentID int      // [NULLABLE] ID to users assignment
	OldValue         string   // [NULLABLE] Value before changing name/email
	NewValue         string   // [NULLABLE] Value after changing name/email
}

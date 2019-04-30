package model

// Review struct
type Review struct {
	ID     int  `json:"id" db:"id"`
	FormID int  `json:"-" db:"form_id"`
	Form   Form `json:"form"`
}

// PeerReview struct
type PeerReview struct {
	ID           int
	ReviewerID   int    // User that is doing the review
	TargetID     int    // User that is getting the review
	ReviewerName string // User that is doing the review
	TargetName   string // User that is getting the review
	AssignmentID int
}

// FullReview holds specific data about an review for displaying it
type FullReview struct {
	Reviewer     int // User that is doing the review
	Target       int // User that is getting the review
	ReviewID     int
	AssignmentID int
	Answers      []ReviewAnswer
}

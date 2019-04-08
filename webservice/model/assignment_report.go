package model

type RawUserReport struct {
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	ReviewsDone  int         `json:"reviews_done"`
	ReviewScores [][]float64 `json:"review_scores"`
}

type ProcessedReviewItem struct {
	Mean   float64 `json:"mean"`
	StdDev float64 `json:"std_dev"`
}

type ProcessedUserReport struct {
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	ReviewsDone int                   `json:"reviews_done"`
	ReviewItems []ProcessedReviewItem `json:"review_items"`
}

type ProcessedAssignmentReport struct {
	UserReports []ProcessedUserReport `json:"user_reports"`

}
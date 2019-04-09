package model

import "github.com/JohanAanesen/CSAMS/webservice/shared/util"

// RawUserReport struct
type RawUserReport struct {
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	ReviewsDone  int         `json:"reviews_done"`
	ReviewScores [][]float64 `json:"review_scores"`
}

// ProcessedReviewItem struct
type ProcessedReviewItem struct {
	Mean   float64 `json:"mean"`
	StdDev float64 `json:"std_dev"`
}

// ProcessedUserReport struct
type ProcessedUserReport struct {
	Name         string                `json:"name"`
	Email        string                `json:"email"`
	ReviewsDone  int                   `json:"reviews_done"`
	ReviewItems  []ProcessedReviewItem `json:"review_items"`
	ReviewMark   float64               `json:"review_mark"`
	ReviewStdDev float64               `json:"review_std_dev"`
}

// ProcessedAssignmentReport struct
type ProcessedAssignmentReport struct {
	UserReports []ProcessedUserReport `json:"user_reports"`
}

// Process raw user report to processed user report
func (raw *RawUserReport) Process() (*ProcessedUserReport, error) {
	// Create struct with base data
	result := ProcessedUserReport{
		Name:        raw.Name,
		Email:       raw.Email,
		ReviewsDone: raw.ReviewsDone,
	}
	// Get the review scores
	scores := raw.ReviewScores
	// Check if slice is bigger then 0
	if len(scores) > 0 {
		// Loop through all slices
		for i := 0; i < len(scores[0]); i++ {
			// Create empty float slice
			data := make([]float64, 0)
			// Loop trough all slices and append it's data
			for j := range scores {
				data = append(data, scores[j][i])
			}
			// Create statistics object
			stats := util.Statistics{
				Entries: data,
			}
			// Calculate mean and standard deviation
			mean, _ := stats.Average()
			stdDev, _ := stats.StandardDeviation()
			// Put data into a struct
			t := ProcessedReviewItem{
				Mean:   mean,
				StdDev: stdDev,
			}
			// Append to the result
			result.ReviewItems = append(result.ReviewItems, t)
		}
	}
	// Return result
	return &result, nil
}

// ExportCSV func
func (par *ProcessedAssignmentReport) ExportCSV() (string, error) {
	return "", nil
}

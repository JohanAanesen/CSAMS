package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"strconv"
)

// ReviewAnswerService struct
type ReviewAnswerService struct {
	reviewAnswerRepo *repository.ReviewAnswerRepository
	courseRepo       *repository.CourseRepository
	assignmentRepo   *repository.AssignmentRepository
	userRepo         *repository.UserRepository
}

// NewReviewAnswerService func
func NewReviewAnswerService(db *sql.DB) *ReviewAnswerService {
	return &ReviewAnswerService{
		reviewAnswerRepo: repository.NewReviewAnswerRepository(db),
		courseRepo:       repository.NewCourseRepository(db),
		assignmentRepo:   repository.NewAssignmentRepository(db),
		userRepo:         repository.NewUserRepository(db),
	}
}

// FetchForAssignment func
func (s *ReviewAnswerService) FetchForAssignment(assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForAssignment(assignmentID)
}

// FetchForTarget func
func (s *ReviewAnswerService) FetchForTarget(target, assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForTarget(target, assignmentID)
}

// FetchForReviewer func
func (s *ReviewAnswerService) FetchForReviewer(reviewer, assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForReviewer(reviewer, assignmentID)
}

// FetchForUser func
func (s *ReviewAnswerService) FetchForUser(userID, assignmentID int) ([][]*model.ReviewAnswer, error) {
	result := make([][]*model.ReviewAnswer, 0)

	reviewers, err := s.FetchReviewUsers(userID, assignmentID)
	if err != nil {
		return result, err
	}

	for _, reviewerID := range reviewers {
		review, err := s.FetchForReviewerAndTarget(reviewerID, userID, assignmentID)
		if err != nil {
			return result, err
		}

		result = append(result, review)
	}

	return result, err
}

// FetchReviewUsers func
func (s *ReviewAnswerService) FetchReviewUsers(target, assignmentID int) ([]int, error) {
	users := make([]int, 0)

	answers, err := s.FetchForTarget(target, assignmentID)
	if err != nil {
		return users, err
	}

	for _, answer := range answers {
		if !util.Contains(users, answer.UserReviewer) {
			users = append(users, answer.UserReviewer)
		}
	}

	return users, err
}

// FetchForReviewerAndTarget func
func (s *ReviewAnswerService) FetchForReviewerAndTarget(reviewer, target, assignmentID int) ([]*model.ReviewAnswer, error) {
	return s.reviewAnswerRepo.FetchForReviewerAndTarget(reviewer, target, assignmentID)
}

// HasBeenReviewed func
func (s *ReviewAnswerService) HasBeenReviewed(target, reviewer, assignmentID int) (bool, error) {
	temp, err := s.reviewAnswerRepo.FetchForReviewerAndTarget(reviewer, target, assignmentID)
	if err != nil {
		return false, err
	}

	return len(temp) > 0, err
}

// CountReviewsDone func
func (s *ReviewAnswerService) CountReviewsDone(userID, assignmentID int) (int, error) {
	return s.reviewAnswerRepo.CountReviewsDone(userID, assignmentID)
}

// Insert func
func (s *ReviewAnswerService) Insert(answer model.ReviewAnswer) (int, error) {
	return s.reviewAnswerRepo.Insert(answer)
}

// FetchMaxScoreFromAssignment func
func (s *ReviewAnswerService) FetchMaxScoreFromAssignment(assignmentID int) (int, error) {
	return s.reviewAnswerRepo.MaxScore(assignmentID)
}

// FetchStatisticsForAssignment func
func (s *ReviewAnswerService) FetchStatisticsForAssignment(assignmentID int) (*util.Statistics, error) {
	// Get max score from assignment
	absMax, err := s.FetchMaxScoreFromAssignment(assignmentID)
	if err != nil {
		return nil, err
	}
	// Create new statistics object
	var result = util.NewStatistics(0, float64(absMax))

	// Get assignment
	assignment, err := s.assignmentRepo.Fetch(assignmentID)
	if err != nil {
		return nil, err
	}
	// Get users from course
	users, err := s.userRepo.FetchAllStudentsFromCourse(assignment.CourseID)
	if err != nil {
		return nil, err
	}
	// Loop through users
	for _, user := range users {
		// Fetch reviews for the user
		reviews, err := s.FetchForUser(user.ID, assignment.ID)
		if err != nil {
			return nil, err
		}
		// Loop through reviews for the user
		for _, review := range reviews {
			// Add result from review
			result.AddEntry(getScoreFromReview(review))
		}
	}

	return result, nil
}

// FetchStatisticsForAssignmentAndUser func
func (s *ReviewAnswerService) FetchStatisticsForAssignmentAndUser(assignmentID, userID int) (*util.Statistics, error) {
	// Get max score from assignment
	absMax, err := s.FetchMaxScoreFromAssignment(assignmentID)
	if err != nil {
		return nil, err
	}
	// Create new statistics object
	var result = util.NewStatistics(0, float64(absMax))

	// Get assignment
	assignment, err := s.assignmentRepo.Fetch(assignmentID)
	if err != nil {
		return nil, err
	}
	// Fetch reviews for the user
	reviews, err := s.FetchForUser(userID, assignment.ID)
	if err != nil {
		return nil, err
	}
	// Loop through reviews for the user
	for _, review := range reviews {
		// Add result from review
		result.AddEntry(getScoreFromReview(review))
	}

	return result, nil
}

// FetchScoreFromReview func
func (s *ReviewAnswerService) FetchScoreFromReview(assignmentID, userID int) ([]float64, error) {
	result := make([]float64, 0)

	reviews, err := s.FetchForUser(userID, assignmentID)
	if err != nil {
		return result, err
	}
	// Loop through reviews for the user
	for _, review := range reviews {
		// Add result from review
		result = append(result, getScoreFromReview(review))
	}

	return result, err
}

// FetchUserReportsForAssignment func
func (s *ReviewAnswerService) FetchUserReportsForAssignment(assignmentID int) ([]model.RawUserReport, error) {
	assignment, err := s.assignmentRepo.Fetch(assignmentID)
	if err != nil {
		return nil, err
	}

	users, err := s.userRepo.FetchAllStudentsFromCourse(assignment.CourseID)
	if err != nil {
		return nil, err
	}

	userReports := make([]model.RawUserReport, 0)

	// Loop through all users
	for _, user := range users {
		// Fetch reviewers for current user
		reviewers, err := s.FetchReviewUsers(user.ID, assignment.ID)
		if err != nil {
			return nil, err
		}

		tempUserReport := model.RawUserReport{}
		tempUserReport.Name = user.Name
		tempUserReport.Email = user.EmailStudent
		tempUserReport.ReviewsDone, err = s.reviewAnswerRepo.CountReviewsDone(user.ID, assignment.ID)
		if err != nil {
			return nil, err
		}

		// Loop through all reviewers
		for _, reviewerID := range reviewers {
			// Fetch all answers from reviewer to target
			reviewAnswers, err := s.reviewAnswerRepo.FetchForReviewerAndTarget(reviewerID, user.ID, assignment.ID)
			if err != nil {
				return nil, err
			}
			// Declare data-slice
			var data = make([]float64, 0)
			// Loop trough all answers
			for _, answer := range reviewAnswers {
				if answer.Weight != 0 {
					// Check answer type
					data = append(data, getWeight(answer))
				}
			}

			tempUserReport.ReviewScores = append(tempUserReport.ReviewScores, data)
		}

		userReports = append(userReports, tempUserReport)
	}

	return userReports, nil
}

// Update review answer and comment
func (s *ReviewAnswerService) Update(targetID, reviewerID, assignmentID int, answer model.ReviewAnswer) error {
	return s.reviewAnswerRepo.Update(targetID, reviewerID, assignmentID, answer)
}

func getWeight(review *model.ReviewAnswer) float64 {
	switch review.Type {
	case "checkbox":
		if review.Answer == "on" {
			return float64(review.Weight)
		}

	case "radio":
		for k := range review.Choices {
			ans, _ := strconv.Atoi(review.Answer)
			if ans == (k + 1) {
				K := float64(k + 1)
				L := float64(len(review.Choices))
				V := float64(review.Weight) * (K / L)
				return V
			}
		}
	}

	return 0
}

func getScoreFromReview(review []*model.ReviewAnswer) float64 {
	var score float64

	for _, item := range review {
		score += getWeight(item)
	}

	return score
}

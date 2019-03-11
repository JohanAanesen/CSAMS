package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
)

// Review struct
type Review struct {
	ID     int  `json:"id" db:"id"`
	FormID int  `json:"-" db:"form_id"`
	Form   Form `json:"form"`
}

// ReviewAnswer holds the data for a single review answer
type ReviewAnswer struct {
	Type   string
	Name   string
	Label  string
	Answer string
}

// Full review holds specific data about an review for displaying it
type FullReview struct {
	Reviewer     int // User that is doing the review
	Target       int // User that is getting the review
	ReviewID     int
	AssignmentID int
	Answers      []ReviewAnswer
}

// ReviewRepository struct handles all database action for review-forms
type ReviewRepository struct{}

// Insert a form into the database
func (repo *ReviewRepository) Insert(form Form) error {
	// Declare FormRepository
	var formRepo = FormRepository{}

	// Insert form to database, and get last inserted Id from it
	formID, err := formRepo.Insert(form)
	if err != nil {
		return err
	}

	// Insertions query
	query := "INSERT INTO reviews (form_id) VALUES(?)"
	// Insert form_id into submissions
	rows, err := db.GetDB().Query(query, formID)
	// Check for error
	if err != nil {
		return err
	}
	// Close the connections
	defer rows.Close()

	// Loop trough fields in the forms
	for _, field := range form.Fields {
		// Insertion query
		query := "INSERT INTO fields (form_id, type, name, label, description, priority, weight, choices) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		// Execute the query
		rows, err := db.GetDB().Query(query, int(formID), field.Type, field.Name, field.Label, field.Description, field.Order, field.Weight, field.Choices)
		// Check for error
		if err != nil {
			return err
		}
		// Close the connection
		rows.Close()
	}

	// Return no error
	return nil
}

// GetAll returns all submission in the database
func (repo *ReviewRepository) GetAll() ([]Review, error) {
	// Declare return slice
	var result []Review
	// Create query-string
	query := "SELECT id, form_id FROM reviews"
	// Perform query
	rows, err := db.GetDB().Query(query)
	// Check for error
	if err != nil {
		return result, err
	}
	// Close connection
	defer rows.Close()

	// Loop through rows
	for rows.Next() {
		// Declare a single Submission
		var review = Review{}
		// Scan the data from the rows
		err = rows.Scan(&review.ID, &review.FormID)
		// Check for error
		if err != nil {
			return result, err
		}

		// Append scan-result to result
		result = append(result, review)
	}

	// Declare a FormRepository
	var formRepo = FormRepository{}
	// Loop through all submissions
	for index, review := range result {
		// Get form from database
		form, err := formRepo.Get(review.FormID)
		// Check for error
		if err != nil {
			return result, nil
		}
		// Get the form
		review.Form = form
		// Set the new values
		result[index] = review
	}

	return result, nil
}

// Update a form in the database
// Deletes all fields, and recreates them
func (repo *ReviewRepository) Update(form Form) error {
	query := "UPDATE forms SET prefix=?, name=?, description=? WHERE id=?"
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	_, err = db.GetDB().Exec(query, form.Prefix, form.Name, form.Description, form.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	query = "DELETE FROM fields WHERE form_id=?"
	rows, err := db.GetDB().Query(query, form.ID)
	if err != nil {
		return err
	}
	rows.Close()

	// Loop trough fields in the forms
	for _, field := range form.Fields {
		// Insertion query
		query := "INSERT INTO fields (form_id, type, name, label, description, priority, weight, choices) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		// Execute the query
		rows, err := db.GetDB().Query(query, form.ID, field.Type, field.Name, field.Label, field.Description, field.Order, field.Weight, field.Choices)
		// Check for error
		if err != nil {
			return err
		}
		// Close the connection
		rows.Close()
	}

	// Return no error
	return nil
}

// Get single review form from the database
func (repo *ReviewRepository) GetSingle(assignmentID int) (Review, error) {
	result := Review{}

	query := "SELECT review_id FROM assignments WHERE id=?"
	rows, err := db.GetDB().Query(query, assignmentID)
	if err != nil {
		return result, err
	}

	var reviewID int

	for rows.Next() {
		err = rows.Scan(&reviewID)
		if err != nil {
			return result, err
		}
	}

	result.ID = reviewID

	query = "SELECT f.form_id, f.id, f.type, f.name, f.label, f.description, f.priority, f.weight, f.choices FROM fields AS f WHERE f.form_id IN (SELECT s.form_id FROM reviews AS s WHERE id IN (SELECT a.review_id FROM assignments AS a WHERE id=?)) ORDER BY f.priority"
	rows, err = db.GetDB().Query(query, assignmentID)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	var formID int
	var form Form

	for rows.Next() {
		var field Field

		err = rows.Scan(&formID, &field.ID, &field.Type, &field.Name, &field.Label, &field.Description, &field.Order, &field.Weight, &field.Choices)
		if err != nil {
			return result, err
		}

		form.Fields = append(form.Fields, field)
	}

	form.ID = formID

	result.Form = form
	result.FormID = formID

	return result, err
}

// InsertReviewAnswers inserts answers from a review into the database
func (repo *ReviewRepository) InsertReviewAnswers(fr FullReview) error {
	query := "INSERT user_reviews (user_reviewer, user_target, review_id, assignment_id, type, name, label, answer) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	for _, answer := range fr.Answers {
		_, err := db.GetDB().Exec(query, fr.Reviewer, fr.Target, fr.ReviewID, fr.AssignmentID, answer.Type, answer.Name, answer.Label, answer.Answer)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

// GetUserWhoHasDoneReview returns user ID for users who have done the peer-review
func (repo *ReviewRepository) GetUserWhoHasDoneReview(userId, assignmentId int) ([]int, error) {
	result := make([]int, 0)

	query := "SELECT DISTINCT user_reviewer FROM user_reviews WHERE user_target=? AND assignment_id=?"
	rows, err := db.GetDB().Query(query, userId, assignmentId)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var temp int

		err = rows.Scan(&temp)
		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}

	return result, err
}

// GetAnswersFromReview returns answers from an review based on user id (target, reviewer) and assignment id
func (repo *ReviewRepository) GetAnswersFromReview(target, reviewer, assignment int) ([]ReviewAnswer, error) {
	result := make([]ReviewAnswer, 0)
	query := "SELECT type, name, label, answer FROM user_reviews WHERE user_target=? AND user_reviewer=? AND assignment_id=?"
	rows, err := db.GetDB().Query(query, target, reviewer, assignment)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var temp ReviewAnswer

		err = rows.Scan(&temp.Type, &temp.Name, &temp.Label, &temp.Answer)
		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}

	return result, err
}

// HasBeenReviewed checks if a user
func (repo *ReviewRepository) HasBeenReviewed(target, reviewer, assignment int) (bool, error) {
	temp, err := repo.GetAnswersFromReview(target, reviewer, assignment)
	if err != nil {
		return false, err
	}

	if len(temp) > 0 {
		return true, nil
	}

	return false, nil
}

// GetReviewForUser returns answered reviews for a single user at a given assignment
func (repo *ReviewRepository) GetReviewForUser(user, assignment int) ([]FullReview, error) {
	result := make([]FullReview, 0)

	query := "SELECT DISTINCT user_reviewer FROM user_reviews WHERE user_target=? AND assignment_id=?"
	rows, err := db.GetDB().Query(query, user, assignment)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var tempID int

		err = rows.Scan(&tempID)
		if err != nil {
			return result, err
		}

		query = "SELECT type, name, label, answer FROM user_reviews WHERE user_target=? AND assignment_id=? AND user_reviewer=?"
		nextRows, err := db.GetDB().Query(query, user, assignment, tempID)
		if err != nil {
			return result, err
		}

		fullReview := FullReview{
			Reviewer: tempID,
			Answers:  make([]ReviewAnswer, 0),
		}

		for nextRows.Next() {
			var reviewAnswer ReviewAnswer

			err = nextRows.Scan(&reviewAnswer.Type, &reviewAnswer.Name, &reviewAnswer.Label, &reviewAnswer.Answer)
			if err != nil {
				return result, err
			}

			fullReview.Answers = append(fullReview.Answers, reviewAnswer)
		}

		result = append(result, fullReview)
	}

	return result, err
}

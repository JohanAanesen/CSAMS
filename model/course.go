package model

// Courses hold the data for a slice of Course-struct
type Courses struct {
	Items []Course `json:"courses"`
}

// Course holds the data for courses
type Course struct {
	ID          string          `json:"id"`
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Teacher     int          `json:"teacher"`
	Year        string       `json:"year"`
	Semester    string       `json:"semester"`
	Assignments []Assignment `json:"assignments"`
}

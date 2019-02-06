package page

// Courses hold the data for a slice of Course-structs
type Courses struct {
	Items []Course `json:"courses"`
}

// Course holds the data for courses
type Course struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Description string `json:"description"`
	Year string `json:"year"`
	Semester string `json:"semester"`
	Assignments []Assignment `json:"assignments"`
}

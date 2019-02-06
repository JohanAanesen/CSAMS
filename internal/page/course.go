package page

type Courses struct {
	Items []Course `json:"courses"`
}

type Course struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Description string `json:"description"`
	Year string `json:"year"`
	Semester string `json:"semester"`
	Assignments []Assignment `json:"assignments"`
}

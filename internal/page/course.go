package page

type Courses struct {
	Items []Course `json:"courses"`
}

type Course struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Assignments []Assignment `json:"assignments"`
}

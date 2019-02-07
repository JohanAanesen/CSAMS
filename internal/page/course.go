package page

// Courses hold the data for a slice of Course-struct
type Courses struct {
	Items []Course `json:"courses"`
}

// Course holds the data for courses
type Course struct {
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Link1Name	string		 `json:"link1name"`
	Link1		string		 `json:"link1"`
	Link2Name	string		 `json:"link2name"`
	Link2		string		 `json:"link2"`
	Link3Name	string		 `json:"link3name"`
	Link3		string		 `json:"link3"`
	Year        string       `json:"year"`
	Semester    string       `json:"semester"`
	Assignments []Assignment `json:"assignments"`
}

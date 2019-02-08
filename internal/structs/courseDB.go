package structs

// COurseDB is a struct for course from the Database
type CourseDB struct {
	Id          int
	CourseCode  string
	CourseName  string
	Teacher     int
	Description string
	Year        int
	Semester    string
}

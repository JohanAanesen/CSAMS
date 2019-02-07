package structs

// COurseDB is a struct for course from the Database
type CourseDB struct {
	Id         int
	CourseCode string
	CourseName string
	Teacher    int
	Info       string
	Link1      string
	Link2      string
	Link3      string
}

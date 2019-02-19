package model

// Users hold the data for a slice of Course-struct
type Users struct {
	Items []User `json:"users"`
}

//User struct to hold session data
type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	EmailStudent  string `json:"emailstudent"`
	EmailPrivate  string `json:"emailprivate"`
	Teacher       bool   `json:"teacher"`
	Authenticated bool   `json:"authenticated"`
}

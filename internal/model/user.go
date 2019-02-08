package model

//User struct to hold session data
type User struct {
	ID            int
	Name          string
	EmailStudent  string
	EmailPrivate  string
	Teacher       bool
	Authenticated bool
}

package page

//Menu struct
type Menu struct {
	Items []MenuItem
}

//MenuItem struct
type MenuItem struct {
	Name string
	Href string
}
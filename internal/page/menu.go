package page

type Menu struct {
	Items []MenuItem
}

type MenuItem struct {
	Name string
	Href string
}
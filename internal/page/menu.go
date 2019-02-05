package page

type Menu struct {
	Items []MenuItem `json:"items"`
}

type MenuItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
}
package page

// Menu holds the data for all menu items
type Menu struct {
	Items []MenuItem `json:"items"`
}

// MenuItem holds the data for a single menu item
type MenuItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

package page

// Assignment hold the data for a single assignment
type Assignment struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

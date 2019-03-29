package model

// Field struct
type Field struct {
	ID          int      `json:"id"`
	FormID      int      `json:"form_id"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Label       string   `json:"label"`
	HasComment  bool     `json:"hasComment"`
	Order       int      `json:"order"`
	Weight      int      `json:"weight,omitempty"`
	Choices     []string `json:"choices,omitempty"`
	Required    bool     `json:"required"`
}

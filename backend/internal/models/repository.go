package models

// Repository represents a GitHub repository.
type Repository struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Owner       User    `json:"owner"`
	Private     bool    `json:"private"`
	Description *string `json:"description,omitempty"`
	URL         string  `json:"url"`
	IsTemplate  bool    `json:"is_template"`
	Archived    bool    `json:"archived"`
}

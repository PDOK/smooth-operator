package model

// Author represents the author or owner of the service or dataset
type Author struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

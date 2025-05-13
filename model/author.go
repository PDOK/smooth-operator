package model

// Author represents the author or owner of the service or dataset
type Author struct {
	// Name of the author
	// +kubebuilder:validation:MinLength:=1
	Name string `json:"name"`

	// Email of the author
	// +kubebuilder:validation:Format:=email
	Email string `json:"email"`
}

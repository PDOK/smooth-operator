package model

// IngressRouteURLs
// +kubebuilder:validation:MinItems:=1
// +kubebuilder:validation:MaxItems:=30
type IngressRouteURLs []IngressRouteURL

type IngressRouteURL struct {
	URL URL `json:"url"`
}

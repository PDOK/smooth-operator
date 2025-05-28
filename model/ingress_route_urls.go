package model

// IngressRouteURLs
// +kubebuilder:validation:MinItems:=1
// Without maxItems x-kubernetes-validation complains about an exceeded computation budget as the list can grow infinitely in theory
// +kubebuilder:validation:MaxItems:=30
type IngressRouteURLs []IngressRouteURL

type IngressRouteURL struct {
	URL URL `json:"url"`
}

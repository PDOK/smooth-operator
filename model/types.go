package model

import (
	"fmt"
	"strings"
)

// ServiceType ...
type ServiceType string

const (
	// WMS ...
	WMS ServiceType = "WMS"
	// WFS ...
	WFS ServiceType = "WFS"
	// Atom ...
	Atom ServiceType = "Atom"
	// OGCAPI ...
	OGCAPI ServiceType = "OGCAPI"
)

// ServiceTypeLabel returns the value of the pdok.nl/service-type label for a ServiceType
func (t ServiceType) ServiceTypeLabel() string {
	if t == OGCAPI {
		return "ogc"
	}

	return strings.ToLower(string(t))
}

// ParseServiceType parses the pdok.nl/service-type label to a ServiceType
func ParseServiceType(input string) (ServiceType, error) {
	switch strings.ToUpper(input) {
	case "WMS":
		return WMS, nil
	case "WFS":
		return WFS, nil
	case "ATOM":
		return Atom, nil
	case "OGCAPI":
		return OGCAPI, nil
	default:
		return "unknown", fmt.Errorf("could not parse %s as a ServiceType", input)
	}
}

type LifecyclePhase string

const (
	// PreProd ...
	PreProd LifecyclePhase = "preprod"
	// Prod ...
	Prod LifecyclePhase = "prod"
)

func ParseLifecyclePhase(lifecyclePhase string, namespace string) LifecyclePhase {
	lp := strings.ToLower(lifecyclePhase)
	ns := strings.ToLower(namespace)

	switch {
	case ns == "services-preprod":
		return PreProd
	case lp == "preprod":
		return PreProd
	case lp == "pre-prod":
		return PreProd
	case lp == "prod":
		return Prod
	default:
		return Prod // Default to prod when not in a preprod namespace or lifecyclePhase is not specified
	}
}

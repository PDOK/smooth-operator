package validation

import (
	"errors"
	"fmt"
	"net/url"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateBaseURL
// Checks if the baseURL is https and has a path
func ValidateBaseURL(baseURL string) error {
	parsed, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid BaseURL: %w", err)
	}

	if parsed.Scheme != "https" && parsed.Hostname() != "localhost" {
		return errors.New("invalid BaseURL: must use https scheme")
	}

	if len(parsed.Path) <= 1 {
		return errors.New("invalid BaseURL: must have a path")
	}

	return nil
}

type BaseUrlProvider interface {
	GetBaseUrl() string
}

func CheckBaseUrlImmutability(oldProvider BaseUrlProvider, newProvider BaseUrlProvider, allErrs *field.ErrorList) {
	if oldProvider.GetBaseUrl() != newProvider.GetBaseUrl() {
		*allErrs = append(*allErrs, field.Forbidden(
			field.NewPath("spec").Child("service").Child("baseURL"),
			"is immutable",
		))
	}
}

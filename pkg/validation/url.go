package validation

import (
	"errors"
	"fmt"
	"github.com/pdok/smooth-operator/model"
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

func CheckUrlImmutability(oldURL, newURL model.URL, allErrs *field.ErrorList, path *field.Path) {
	if oldURL.URL == nil && newURL.URL == nil {
		return
	}
	if (oldURL.URL == nil && newURL.URL != nil) || (oldURL.URL != nil && newURL.URL == nil) || (*oldURL.URL != *newURL.URL) {
		*allErrs = append(*allErrs, field.Forbidden(
			path,
			"is immutable",
		))
	}
}

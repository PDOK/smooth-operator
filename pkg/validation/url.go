package validation

import (
	"fmt"
	"net/url"
)

// ValidateBaseURL
// Checks if the baseURL is https and has a path
func ValidateBaseURL(baseURL string) error {
	parsed, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid BaseURL: %w", err)
	}

	if parsed.Scheme != "https" {
		return fmt.Errorf("invalid BaseURL: must use https scheme")
	}

	if len(parsed.Path) <= 1 {
		return fmt.Errorf("invalid BaseURL: must have a path")
	}

	return nil
}

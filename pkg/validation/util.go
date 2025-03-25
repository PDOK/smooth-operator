package validation

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// FormatValidationWarning is a helper function so all warnings have the same formating
func FormatValidationWarning(message string, groupVersionKind schema.GroupVersionKind, name string) string {
	return fmt.Sprintf("%s/%s: %s", groupVersionKind, name, message)
}

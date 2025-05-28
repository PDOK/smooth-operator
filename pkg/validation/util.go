package validation

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// AddWarning is a helper function so all warnings have the same formating
func AddWarning(warnings *[]string, path field.Path, message string, groupVersionKind schema.GroupVersionKind, name string) {
	*warnings = append(*warnings, fmt.Sprintf("%s/%s: %s: %s", groupVersionKind, name, path.String(), message))
}

package validation

import (
	"fmt"
	"strings"
)

// ValidateLabelsOnCCreate
// Checks labels on creation
func ValidateLabelsOnCreate(labels map[string]string) error {
	if len(labels) == 0 {
		return fmt.Errorf("labels must not be empty")
	}

	return nil
}

// ValidateLabelsOnUpdate
// Checks if the old and new label set are exactly the same
func ValidateLabelsOnUpdate(oldLabels, newLabels map[string]string) error {
	reasons := make([]string, 0)
	for oldKey, oldValue := range oldLabels {
		newValue, ok := newLabels[oldKey]
		if !ok {
			reasons = append(reasons, fmt.Sprintf("label '%s' removed", oldKey))
		} else if oldValue != newValue {
			reasons = append(reasons, fmt.Sprintf("label '%s' changed from '%s' to '%s'", oldKey, oldValue, newValue))
		}
	}

	for newKey, _ := range newLabels {
		if _, ok := oldLabels[newKey]; !ok {
			reasons = append(reasons, fmt.Sprintf("label '%s' added", newKey))
		}
	}

	if len(reasons) > 0 {
		return fmt.Errorf("labels are immutable. %s", strings.Join(reasons, ", "))
	}

	return nil
}

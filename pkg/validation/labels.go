package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateLabelsOnCCreate
// Checks labels on creation
func ValidateLabelsOnCreate(labels map[string]string) *field.Error {
	if len(labels) == 0 {
		return field.Required(field.NewPath("metadata").Child("labels"), "can't be empty")
	}

	return nil
}

// ValidateLabelsOnUpdate
// Checks if the old and new label set are exactly the same
func ValidateLabelsOnUpdate(oldLabels, newLabels map[string]string, allErrs *field.ErrorList) {
	fieldPath := field.NewPath("metadata").Child("labels")
	for oldKey, oldValue := range oldLabels {
		newValue, ok := newLabels[oldKey]
		if !ok {
			*allErrs = append(*allErrs, field.Required(fieldPath.Child(oldKey), "labels cannot be removed"))
		} else if oldValue != newValue {
			*allErrs = append(*allErrs, field.Invalid(fieldPath.Child(oldKey), newValue, "immutable: should be: "+oldValue))
		}
	}

	for newKey := range newLabels {
		if _, ok := oldLabels[newKey]; !ok {
			*allErrs = append(*allErrs, field.Forbidden(fieldPath.Child(newKey), "new labels cannot be added"))
		}
	}
}

package reloader

import (
	"flag"
	"fmt"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// FlagName is the name of the operator flag that enables the Reloader annotation on Deployments.
	FlagName = "set-reloader-annotations"
	// AutoAnnotation makes Stakater Reloader restart a Deployment when its referenced ConfigMaps or Secrets change.
	AutoAnnotation = "reloader.stakater.com/auto"
	// AutoAnnotationValue is the value set for AutoAnnotation.
	AutoAnnotationValue = "true"
	// SearchAnnotation makes Stakater Reloader restart a Deployment when referenced ConfigMaps or Secrets annotated with reloader.stakater.com/match change.
	SearchAnnotation = "reloader.stakater.com/search"
	// SearchAnnotationValue is the value set for SearchAnnotation.
	SearchAnnotationValue = "true"
)

type annotationMode string

const (
	disabledMode annotationMode = "false"
	autoMode     annotationMode = "auto"
	searchMode   annotationMode = "search"
)

var reloaderMode = disabledMode

func (m *annotationMode) String() string {
	return string(*m)
}

func (m *annotationMode) Set(value string) error {
	switch value {
	case string(autoMode), string(searchMode):
		*m = annotationMode(value)
	default:
		enabled, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("expected auto, search, true or false, got %q", value)
		}
		*m = disabledMode
		if enabled {
			*m = autoMode
		}
	}
	return nil
}

func (*annotationMode) IsBoolFlag() bool {
	return true
}

// BindFlag registers the set-reloader-annotations flag on the given FlagSet.
func BindFlag(fs *flag.FlagSet) {
	reloaderMode = disabledMode
	fs.Var(&reloaderMode, FlagName,
		"Set the Stakater Reloader annotation on Deployments: auto, search, true (auto), or false (disabled). A bare flag enables auto; use =search to enable search.")
}

// SetReloaderAnnotations enables the auto annotation or disables both Reloader annotations.
func SetReloaderAnnotations(set bool) {
	reloaderMode = disabledMode
	if set {
		reloaderMode = autoMode
	}
}

// IsEnabled reports whether the Reloader annotation should be set.
func IsEnabled() bool {
	return reloaderMode != disabledMode
}

// SetAnnotation sets the selected Reloader annotation and removes the unselected annotations.
func SetAnnotation(obj metav1.Object) {
	annotations := obj.GetAnnotations()
	delete(annotations, AutoAnnotation)
	delete(annotations, SearchAnnotation)
	if IsEnabled() {
		if annotations == nil {
			annotations = map[string]string{}
		}
		if reloaderMode == searchMode {
			annotations[SearchAnnotation] = SearchAnnotationValue
		} else {
			annotations[AutoAnnotation] = AutoAnnotationValue
		}
	}
	obj.SetAnnotations(annotations)
}

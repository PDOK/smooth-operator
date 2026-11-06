package reloader

import (
	"flag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// FlagName is the name of the operator flag that enables the Reloader annotation on Deployments.
	FlagName = "set-reloader-annotations"
	// AutoAnnotation makes Stakater Reloader restart a Deployment when its referenced ConfigMaps or Secrets change.
	AutoAnnotation = "reloader.stakater.com/auto"
	// AutoAnnotationValue is the value set for AutoAnnotation.
	AutoAnnotationValue = "true"
)

var setReloaderAnnotations = false

// BindFlag registers the set-reloader-annotations flag on the given FlagSet.
func BindFlag(fs *flag.FlagSet) {
	fs.BoolVar(&setReloaderAnnotations, FlagName, false,
		"When enabled Deployments get the "+AutoAnnotation+" annotation that is used by Stakater Reloader.")
}

// SetReloaderAnnotations enables or disables setting the Reloader annotation.
func SetReloaderAnnotations(set bool) {
	setReloaderAnnotations = set
}

// IsEnabled reports whether the Reloader annotation should be set.
func IsEnabled() bool {
	return setReloaderAnnotations
}

// SetAnnotation adds the Reloader annotation to obj when enabled, and removes it otherwise.
func SetAnnotation(obj metav1.Object) {
	annotations := obj.GetAnnotations()
	if setReloaderAnnotations {
		if annotations == nil {
			annotations = map[string]string{}
		}
		annotations[AutoAnnotation] = AutoAnnotationValue
		obj.SetAnnotations(annotations)
		return
	}
	if _, ok := annotations[AutoAnnotation]; ok {
		delete(annotations, AutoAnnotation)
		obj.SetAnnotations(annotations)
	}
}

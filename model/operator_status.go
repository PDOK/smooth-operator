package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// OperatorStatus defines the observed state of an Atom/WFS/WMS/....
type OperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Each condition contains details for one aspect of the current state of this Atom.
	// Known .status.conditions.type are: "Reconciled"
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// The result of creating or updating of each derived resource for this Atom.
	OperationResults map[string]controllerutil.OperationResult `json:"operationResults,omitempty"`
}

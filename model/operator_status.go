package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type ReplicaSetStatus struct {
	Generation  int32 `json:"generation"`
	Total       int32 `json:"total"`
	Ready       int32 `json:"ready"`
	Available   int32 `json:"available"`
	Unavailable int32 `json:"unavailable"`
}

type PodSummary []ReplicaSetStatus

// OperatorStatus defines the observed state of an Atom/WFS/WMS/....
type OperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PodSummary PodSummary `json:"podSummary,omitempty"`
	// Each condition contains details for one aspect of the current state of this Atom.
	// Known .status.conditions.type are: "Reconciled"
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// The result of creating or updating of each derived resource for this Atom.
	OperationResults map[string]controllerutil.OperationResult `json:"operationResults,omitempty"`
}

package status

import (
	"context"
	"time"

	"github.com/pdok/smooth-operator/model"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	reconciledConditionType          = "Reconciled"
	reconciledConditionReasonSuccess = "Success"
	reconciledConditionReasonError   = "Error"
)

type ObjectWithStatus interface {
	client.Object
	OperatorStatus() *model.OperatorStatus
}

func LogAndUpdateStatusError[O ObjectWithStatus](ctx context.Context, k8sClient client.Client, obj O, err error) {
	lgr := log.FromContext(ctx)
	lgr.Error(err, "reconcile error")

	generation := obj.GetGeneration()
	updateStatus(ctx, k8sClient, obj, []metav1.Condition{{
		Type:               reconciledConditionType,
		Status:             metav1.ConditionFalse,
		Reason:             reconciledConditionReasonError,
		Message:            err.Error(),
		ObservedGeneration: generation,
		LastTransitionTime: metav1.NewTime(time.Now()),
	}}, nil)
}

func LogAndUpdateStatusFinished[O ObjectWithStatus](ctx context.Context, k8sClient client.Client, obj O, operationResults map[string]controllerutil.OperationResult) {
	lgr := log.FromContext(ctx)
	lgr.Info("operation results", "results", operationResults)

	generation := obj.GetGeneration()
	updateStatus(ctx, k8sClient, obj, []metav1.Condition{{
		Type:               reconciledConditionType,
		Status:             metav1.ConditionTrue,
		Reason:             reconciledConditionReasonSuccess,
		ObservedGeneration: generation,
		LastTransitionTime: metav1.NewTime(time.Now()),
	}}, operationResults)
}

func updateStatus[O ObjectWithStatus](ctx context.Context, k8sClient client.Client, obj O, conditions []metav1.Condition, operationResults map[string]controllerutil.OperationResult) {
	lgr := log.FromContext(ctx)
	if err := k8sClient.Get(ctx, client.ObjectKeyFromObject(obj), obj); err != nil {
		log.FromContext(ctx).Error(err, "unable to update status")
		return
	}

	status := obj.OperatorStatus()
	if status == nil {
		status = &model.OperatorStatus{}
	}

	podSummary, err := getPodSummary(ctx, k8sClient, obj)
	if err != nil {
		lgr.Error(err, "unable to get pod summary for status update")
		return
	}

	changed := false
	if !equality.Semantic.DeepEqual(status.PodSummary, podSummary) {
		status.PodSummary = podSummary
		changed = true
	}
	for _, condition := range conditions {
		if meta.SetStatusCondition(&status.Conditions, condition) {
			changed = true
		}
	}
	if !equality.Semantic.DeepEqual(status.OperationResults, operationResults) {
		status.OperationResults = operationResults
		changed = true
	}
	if !changed {
		return
	}
	if err := k8sClient.Status().Update(ctx, obj); err != nil {
		lgr.Error(err, "unable to update status")
	}
}

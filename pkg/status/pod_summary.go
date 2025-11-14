package status

import (
	"context"
	"errors"
	"slices"
	"strconv"

	"github.com/pdok/smooth-operator/model"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// GetReplicaSetEventHandlerForObj returns an event handler that only triggers a reconcile
// if a ReplicaSet is owned by obj through a Deployment, in other words the replicaset is a grandchild of obj
func GetReplicaSetEventHandlerForObj(mgr ctrl.Manager, objKind string) handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
		replicaset, ok := obj.(*appsv1.ReplicaSet)
		if !ok {
			return nil
		}

		return checkReplicaSetBelongsToKind(ctx, mgr, replicaset, objKind)
	})
}

func checkReplicaSetBelongsToKind(ctx context.Context, mgr ctrl.Manager, replicaset *appsv1.ReplicaSet, objKind string) []reconcile.Request {
	if deploymentRef := getOwnerRefOfKind(replicaset, "Deployment"); deploymentRef != nil {
		deployment := &appsv1.Deployment{}
		if err := mgr.GetClient().Get(ctx, types.NamespacedName{
			Name:      deploymentRef.Name,
			Namespace: replicaset.GetNamespace(),
		}, deployment); err != nil {
			return nil
		}

		if objRef := getOwnerRefOfKind(deployment, objKind); objRef != nil {
			log.FromContext(ctx).Info("Reconcile requested", "kind", objKind)
			return []reconcile.Request{{
				NamespacedName: types.NamespacedName{
					Name:      objRef.Name,
					Namespace: deployment.GetNamespace(),
				},
			}}
		}
	}

	return nil
}

func getOwnerRefOfKind(childObj client.Object, kind string) *metav1.OwnerReference {
	for _, owner := range childObj.GetOwnerReferences() {
		if owner.Kind == kind {
			return &owner
		}
	}

	return nil
}

// getPodSummary returns a pod summary that includes the status of the last two replica sets that belong to obj based on its labels
func getPodSummary(ctx context.Context, k8sClient client.Client, obj client.Object) (model.PodSummary, error) {
	var replicaSetList appsv1.ReplicaSetList
	err := k8sClient.List(ctx, &replicaSetList, client.MatchingLabels(obj.GetLabels()))
	if err != nil {
		return nil, err
	}

	replicaSetRevision := func(rs appsv1.ReplicaSet) (int, error) {
		val, ok := rs.Annotations["deployment.kubernetes.io/revision"]
		if !ok {
			return 0, errors.New("annotation 'deployment.kubernetes.io/revision' missing from replicaset")
		}

		return strconv.Atoi(val)
	}

	// Sort replicasets by revision
	replicaSets := replicaSetList.Items
	if len(replicaSets) > 1 {
		slices.SortFunc(replicaSets, func(rsa, rsb appsv1.ReplicaSet) int {
			revA, aErr := replicaSetRevision(rsa)
			if aErr != nil {
				err = aErr
				return 0
			}

			revB, bErr := replicaSetRevision(rsb)
			if bErr != nil {
				err = bErr
				return 0
			}

			// Descending order
			return revB - revA
		})
		if err != nil {
			return nil, err
		}
	}

	ps := model.PodSummary{}
	for i, rs := range replicaSets {
		revision, err := replicaSetRevision(rs)
		if err != nil {
			return nil, err
		}

		ps = append(ps, model.ReplicaSetStatus{
			//nolint:gosec
			Generation:  int32(revision),
			Total:       rs.Status.Replicas,
			Ready:       rs.Status.ReadyReplicas,
			Available:   rs.Status.AvailableReplicas,
			Unavailable: rs.Status.Replicas - rs.Status.ReadyReplicas,
		})

		// Stop after two
		if i == 1 {
			break
		}
	}

	return ps, nil
}

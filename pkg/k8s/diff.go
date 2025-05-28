package k8s

import (
	"context"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func ShowDiff(ctx context.Context, k8sClient client.Client, obj client.Object, f controllerutil.MutateFn) {
	lgr := log.FromContext(ctx)

	err := k8sClient.Get(ctx, client.ObjectKeyFromObject(obj), obj)
	if err != nil {
		// Object not found, not able to show diff
		return
	}

	current := obj.DeepCopyObject().(client.Object)
	_ = f()

	if diff := cmp.Diff(current, obj); diff != "" {
		lgr.Info(diff)
	}
}

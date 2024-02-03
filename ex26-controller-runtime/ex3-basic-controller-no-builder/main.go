package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const LabelPodsCount = "570499536.xyz/pods-count"

func main() {
	logf.SetLogger(zap.New())
	log := logf.Log.WithName("ex1-basic-controller-using-builder")

	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{
		Cache: cache.Options{
			DefaultNamespaces: map[string]cache.Config{
				"default": cache.Config{},
			},
		},
	})
	if err != nil {
		log.Error(err, "cannot create a manager")
		os.Exit(1)
	}

	client, err := controller.New("ReplicaSet", mgr, controller.Options{
		Reconciler: &ReplicaSetReconciler{
			Log:    log.WithName("ReplicaSetReconciler"),
			Client: mgr.GetClient(),
		},
	})
	if err != nil {
		log.Error(err, "cannot create the controller")
		os.Exit(1)
	}
	err = client.Watch(
		source.Kind(mgr.GetCache(), &appsv1.ReplicaSet{}),
		&handler.EnqueueRequestForObject{},
	)
	if err != nil {
		log.Error(err, "cannot watch ReplicaSet")
		os.Exit(1)
	}
	err = client.Watch(
		source.Kind(mgr.GetCache(), &corev1.Pod{}),
		handler.EnqueueRequestForOwner(
			mgr.GetScheme(),
			mgr.GetRESTMapper(),
			&appsv1.ReplicaSet{},
			handler.OnlyControllerOwner(),
		),
	)
	if err != nil {
		log.Error(err, "cannot watch Pod")
		os.Exit(1)
	}

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "cannot start the manager")
		os.Exit(1)
	}
}

type ReplicaSetReconciler struct {
	Log logr.Logger
	client.Client
}

func (r *ReplicaSetReconciler) Reconcile(
	ctx context.Context,
	req reconcile.Request,
) (reconcile.Result, error) {
	replicaSet := &appsv1.ReplicaSet{}
	if err := r.Get(ctx, req.NamespacedName, replicaSet); err != nil {
		r.Log.Error(err, "cannot get ReplicaSet", "NamespacedName", req.NamespacedName)
		return reconcile.Result{}, err
	}

	podList := &corev1.PodList{}
	if err := r.List(ctx, podList,
		client.InNamespace(req.Namespace),
		client.MatchingLabels(replicaSet.Spec.Selector.MatchLabels)); err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	if replicaSet.Labels == nil {
		replicaSet.Labels = make(map[string]string)
	}
	replicaSet.Labels[LabelPodsCount] = strconv.Itoa(len(podList.Items))

	if err := r.Update(ctx, replicaSet); err != nil {
		return reconcile.Result{}, err
	}

	r.Log.Info(
		"add label",
		"replicaSet",
		req.NamespacedName,
		"label",
		fmt.Sprintf("%s=%d", LabelPodsCount, len(podList.Items)),
	)
	return reconcile.Result{}, nil
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	gamev1alpha1 "github.com/uhziel/game-operator/api/v1alpha1"
)

const LabelPodsCount = "570499536.xyz/pods-count"

var Scheme = runtime.NewScheme()

func init() {
	clientgoscheme.AddToScheme(Scheme)
	gamev1alpha1.AddToScheme(Scheme)
}

func main() {
	logf.SetLogger(zap.New())
	log := logf.Log.WithName("ex4-crd-controller-using-builder")

	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{
		Scheme: Scheme,
	})
	if err != nil {
		log.Error(err, "cannot create a manager")
		os.Exit(1)
	}

	if err := builder.ControllerManagedBy(mgr).
		For(&gamev1alpha1.Server{}).
		Complete(&ServerReconciler{
			Log:    log.WithName("ServerReconciler"),
			Client: mgr.GetClient(),
		}); err != nil {
		log.Error(err, "cannot create the controller")
		os.Exit(1)
	}

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "cannot start the manager")
		os.Exit(1)
	}
}

type ServerReconciler struct {
	Log logr.Logger
	client.Client
}

func (r *ServerReconciler) Reconcile(
	ctx context.Context,
	req reconcile.Request,
) (reconcile.Result, error) {

	server := &gamev1alpha1.Server{}
	if err := r.Get(ctx, req.NamespacedName, server); err != nil {
		return reconcile.Result{}, err
	}

	if server.Labels == nil {
		server.Labels = make(map[string]string)
	}
	server.Labels[LabelPodsCount] = "1"

	if err := r.Update(ctx, server); err != nil {
		return reconcile.Result{}, err
	}

	r.Log.Info(
		"add label",
		"server",
		req.NamespacedName,
		"label",
		fmt.Sprintf("%s=%d", LabelPodsCount, 1),
	)
	return reconcile.Result{}, nil
}

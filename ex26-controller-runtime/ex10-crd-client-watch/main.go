package main

import (
	"os"

	gamev1alpha1 "github.com/uhziel/game-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var Scheme = runtime.NewScheme()

func init() {
	clientgoscheme.AddToScheme(Scheme)
	gamev1alpha1.AddToScheme(Scheme)
}

func main() {
	logf.SetLogger(zap.New())
	log := logf.Log.WithName("ex5-crd-client")

	c, err := client.NewWithWatch(config.GetConfigOrDie(), client.Options{
		Scheme: Scheme,
	})

	if err != nil {
		log.Error(err, "cannot create client")
		os.Exit(1)
	}

	ctx := signals.SetupSignalHandler()
	serverList := &gamev1alpha1.ServerList{}
	if err := c.List(ctx, serverList, client.InNamespace("default")); err != nil {
		log.Error(err, "cannot list ServerList")
		os.Exit(1)
	}

	for _, server := range serverList.Items {
		log.Info("server", "name", server.Name)
	}

	// watch
	watchInterface, err := c.Watch(ctx, &gamev1alpha1.ServerList{}, &client.ListOptions{
		Namespace:     "test",
		FieldSelector: fields.OneTermEqualSelector("metadata.name", "test"),
	})
	if err != nil {
		panic(err)
	}

	go func() {
		defer watchInterface.Stop()
		for e := range watchInterface.ResultChan() {
			server := e.Object.(*gamev1alpha1.Server)
			log.Info(
				"watched event",
				"type",
				e.Type,
				"status",
				server.Status.PrintableStatus,
				"addr",
				server.Status.Address,
			)
		}
	}()

	select {
	case <-ctx.Done():
		break
	}
}

package main

import (
	"context"
	"os"

	gamev1alpha1 "github.com/uhziel/game-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var Scheme = runtime.NewScheme()

func init() {
	clientgoscheme.AddToScheme(Scheme)
	gamev1alpha1.AddToScheme(Scheme)
}

func main() {
	logf.SetLogger(zap.New())
	log := logf.Log.WithName("ex5-crd-client")

	c, err := client.New(config.GetConfigOrDie(), client.Options{
		Scheme: Scheme,
	})

	if err != nil {
		log.Error(err, "cannot create client")
		os.Exit(1)
	}

	ctx := context.Background()
	serverList := &gamev1alpha1.ServerList{}
	if err := c.List(ctx, serverList, client.InNamespace("default")); err != nil {
		log.Error(err, "cannot list ServerList")
		os.Exit(1)
	}

	for _, server := range serverList.Items {
		log.Info("server", "name", server.Name)
	}
}

package main

import (
	"context"
	"os"

	gamev1alpha1 "github.com/uhziel/game-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
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
	log := logf.Log.WithName("ex7-crd-client-with-cache-fieldindexer")

	cfg := config.GetConfigOrDie()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cl, err := cluster.New(cfg, func(o *cluster.Options) {
		o.Scheme = Scheme
		o.Cache.DefaultNamespaces = map[string]cache.Config{
			"default": cache.Config{},
		}
	})
	if err != nil {
		log.Error(err, "cannot create cluster")
		os.Exit(1)
	}

	cl.GetFieldIndexer().IndexField(
		ctx,
		&gamev1alpha1.Server{},
		".zhulei.status",
		func(o client.Object) []string {
			res := []string{}

			server := o.(*gamev1alpha1.Server)
			res = append(res, string(server.Status.PrintableStatus))

			return res
		},
	)

	go func() {
		cl.Start(ctx)
	}()

	if !cl.GetCache().WaitForCacheSync(ctx) {
		log.Info("cannot WaitForCacheSync")
		os.Exit(1)
	}

	serverList := &gamev1alpha1.ServerList{}
	if err := cl.GetClient().List(ctx, serverList, client.InNamespace("default"), client.MatchingFields(map[string]string{".zhulei.status": "Offline"})); err != nil {
		log.Error(err, "cannot list ServerList")
		os.Exit(1)
	}

	for _, server := range serverList.Items {
		log.Info("server", "name", server.Name)
	}
}

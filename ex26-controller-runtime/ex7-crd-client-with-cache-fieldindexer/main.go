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

	clientCache, err := cache.New(cfg, cache.Options{
		Scheme: Scheme,
		DefaultNamespaces: map[string]cache.Config{
			"default": cache.Config{},
		},
	})
	if err != nil {
		log.Error(err, "cannot create cache")
		os.Exit(1)
	}

	clientCache.IndexField(
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
		clientCache.Start(ctx)
	}()

	c, err := client.New(cfg, client.Options{
		Scheme: Scheme,
		Cache: &client.CacheOptions{
			Reader: clientCache,
		},
	})

	if err != nil {
		log.Error(err, "cannot create client")
		os.Exit(1)
	}

	if !clientCache.WaitForCacheSync(ctx) {
		log.Error(err, "cannot WaitForCacheSync")
		os.Exit(1)
	}

	serverList := &gamev1alpha1.ServerList{}
	if err := c.List(ctx, serverList, client.InNamespace("default"), client.MatchingFields(map[string]string{".zhulei.status": "Offline"})); err != nil {
		log.Error(err, "cannot list ServerList")
		os.Exit(1)
	}

	for _, server := range serverList.Items {
		log.Info("server", "name", server.Name)
	}
}

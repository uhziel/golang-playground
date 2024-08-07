package main

import (
	"fmt"
	"os"

	"github.com/go-logr/logr"
	gamev1alpha1 "github.com/uhziel/game-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var Scheme = runtime.NewScheme()

func init() {
	clientgoscheme.AddToScheme(Scheme)
	gamev1alpha1.AddToScheme(Scheme)
}

// isInInitialList 为 false 则表示 watch 期间新创建的；否则，来自已有列表。
// 在添加 eventHandler 时，已有列表中的对象会被发送 (added, isInInitialList=true)，是为了避免出现遗漏。
func eventHandler(log logr.Logger) toolscache.ResourceEventHandlerDetailedFuncs {
	return toolscache.ResourceEventHandlerDetailedFuncs{
		AddFunc: func(obj interface{}, isInInitialList bool) {
			serverInstance := obj.(*gamev1alpha1.ServerInstance)
			log.Info(
				"ServerInstance added",
				"Name",
				serverInstance.Name,
				"isInInitialList",
				isInInitialList,
			)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			serverInstance := newObj.(*gamev1alpha1.ServerInstance)
			log.Info("ServerInstance updated", "Name", serverInstance.Name)
		},
		DeleteFunc: func(obj interface{}) {
			serverInstance := obj.(*gamev1alpha1.ServerInstance)
			log.Info("ServerInstance deleted", "Name", serverInstance.Name)
		},
	}
}

func prompt() {
	fmt.Println("> Press Enter to continue")
	fmt.Scanln()
}

func main() {
	logf.SetLogger(zap.New())
	log := logf.Log.WithName("ex9-crd-cluster-watch")

	cfg := config.GetConfigOrDie()

	ctx := signals.SetupSignalHandler()

	cl, err := cluster.New(cfg, func(o *cluster.Options) {
		o.Scheme = Scheme
		//o.Cache.DefaultNamespaces = map[string]cache.Config{
		//	"default": {},
		//}
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

	serverInstanceInformer, err := cl.GetCache().GetInformer(ctx, &gamev1alpha1.ServerInstance{})
	if err != nil {
		log.Error(err, "GetInformer fail", "Kind", "ServerInstance")
		os.Exit(1)
	}

	if _, err := serverInstanceInformer.AddEventHandler(eventHandler(log)); err != nil {
		log.Error(err, "AddEventHandler fail", "Kind", "ServerInstance")
		os.Exit(1)
	}

	go func() {
		cl.Start(ctx)
	}()

	if !cl.GetCache().WaitForCacheSync(ctx) {
		log.Info("cannot WaitForCacheSync")
		os.Exit(1)
	}

	log.Info("started")

	serverList := &gamev1alpha1.ServerList{}
	if err := cl.GetClient().List(ctx, serverList, client.InNamespace("default"), client.MatchingFields(map[string]string{".zhulei.status": "Offline"})); err != nil {
		log.Error(err, "cannot list ServerList")
		os.Exit(1)
	}

	for _, server := range serverList.Items {
		log.Info("server", "name", server.Name)
	}

	prompt()
	fmt.Println("adding eventHandler2")
	// 可以在运行后添加消息处理函数，已有的对象会触发 (added, isInInitialList=true)
	if _, err := serverInstanceInformer.AddEventHandler(eventHandler(log.WithName("eventHandler2"))); err != nil {
		log.Error(err, "AddEventHandler fail", "Kind", "ServerInstance")
		os.Exit(1)
	}

	select {
	case <-ctx.Done():
		break
	}
}

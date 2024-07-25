package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache" // 最重要的包
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue" // 最重要的包
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func connectToK8s(kubeconfig string) *kubernetes.Clientset {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func createJob(clientset *kubernetes.Clientset, jobName, imageName string) {
	jobs := clientset.BatchV1().Jobs("default")
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  jobName,
							Image: imageName,
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("Failed to create K8s job.")
	}

	//print job details
	log.Println("Created K8s job successfully")
}

func createComponents(
	clientset *kubernetes.Clientset,
) (cache.Indexer, workqueue.RateLimitingInterface, cache.Controller) {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultItemBasedRateLimiter())

	podListWatcher := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"pods",
		metav1.NamespaceDefault,
		fields.Everything(),
	)
	indexer, informer := cache.NewIndexerInformer(
		podListWatcher,
		&corev1.Pod{},
		0,
		cache.ResourceEventHandlerDetailedFuncs{
			AddFunc: func(obj interface{}, _ bool) {
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err != nil {
					log.Println(err)
				}
				queue.Add(key)
			},
			UpdateFunc: func(_, newObj interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(newObj)
				if err != nil {
					log.Println(err)
				}
				queue.Add(key)
			},
			// OnDelete will get the final state of the item if it is known, otherwise it will get an object of type DeletedFinalStateUnknown. This can happen if the watch is closed and misses the delete event and we don't notice the deletion until the subsequent re-list.
			DeleteFunc: func(obj interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
				if err != nil {
					log.Println(err)
				}
				queue.Add(key)
			},
		},
		cache.Indexers{},
	)

	return indexer, queue, informer
}

type Controller struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller
}

func (c *Controller) Run(numWorkers int, ctx context.Context) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown() // 对应的 c.queue.Get() 的返回值 shutdown 会为 true

	log.Println("starting Controller")

	go c.informer.Run(ctx.Done())
	if !cache.WaitForCacheSync(ctx.Done(), c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("WaitForCacheSync time out "))
		return
	}

	for i := 0; i < numWorkers; i++ {
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	<-ctx.Done()
	log.Println("stopping Controller")
}

func (c *Controller) runWorker(ctx context.Context) {
	for c.processNextItem() {

	}
}

func (c *Controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(item)

	key := item.(string)
	if err := c.reconcile(key); err != nil {
		c.handleErr(err, key)
	} else {
		c.queue.Forget(item)
	}

	return true
}

func (c *Controller) reconcile(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		return err
	}

	if !exists {
		log.Println("cannot find pod ", key)
	} else {
		log.Printf("sync/add/update/delete pod %s\n", obj.(*corev1.Pod).Name)
	}
	return nil
}

func (c *Controller) handleErr(err error, item any) {
	if c.queue.NumRequeues(item) < 3 {
		c.queue.AddRateLimited(item)
		return
	}

	c.queue.Forget(item)
	runtime.HandleError(fmt.Errorf("cannot handle err=%v item=%v", err, item))
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String(
			"kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file",
		)
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	clientset := connectToK8s(*kubeconfig)
	indexer, queue, informer := createComponents(clientset)

	controller := &Controller{
		indexer:  indexer,
		queue:    queue,
		informer: informer,
	}

	ctx := signals.SetupSignalHandler()

	go controller.Run(1, ctx)

	indexer.Add(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "hello",
		},
	})

	select {
	case <-ctx.Done():
		return
	}
}

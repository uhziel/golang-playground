package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
)

func connectToK8s(kubeconfig string) *dynamic.DynamicClient {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return client
}

func prompt() {
	fmt.Println("> Press Enter to continue")
	fmt.Scanln()
}

const (
	deployName = "foobar"
)

func createDeploy(ctx context.Context, deployClient dynamic.ResourceInterface) error {
	deploy := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"name": deployName,
			},
			"spec": map[string]interface{}{
				"replicas": 2, // NOTICE
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": deployName,
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": deployName,
						},
					},
					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  "nginx",
								"image": "nginx:alpine",
							},
						},
					},
				},
			},
		},
	}

	// 注意, newDeploy 是一个被创建后实际被存储到 etcd 的 object，deploy 中没指定的值在 newDeploy 会被填充实际值。
	newDeploy, err := deployClient.Create(ctx, deploy, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create deploy: %w", err)
	}

	fmt.Printf("%s\n", newDeploy.GetName())

	return nil
}

func updateDeploy(ctx context.Context, deployClient dynamic.ResourceInterface) error {
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		deploy, err := getDeploy(ctx, deployClient)
		if err != nil {
			return fmt.Errorf("updateDeploy fail: %w", err)
		}

		unstructured.SetNestedField(
			deploy.Object,
			int64(3),
			"spec",
			"replicas",
		) // NOTICE 直接写 3 会报错"cannot deep copy int"

		_, err = deployClient.Update(ctx, deploy, metav1.UpdateOptions{})
		return err
	})
	return err
}

func deleteDeploy(ctx context.Context, deployClient dynamic.ResourceInterface) error {
	return deployClient.Delete(ctx, deployName, metav1.DeleteOptions{})
}

func listDeploy(
	ctx context.Context,
	deployClient dynamic.ResourceInterface,
) (*unstructured.UnstructuredList, error) {
	return deployClient.List(ctx, metav1.ListOptions{})
}

func getDeploy(
	ctx context.Context,
	deployClient dynamic.ResourceInterface,
) (*unstructured.Unstructured, error) {
	return deployClient.Get(ctx, deployName, metav1.GetOptions{})
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

	ctx := context.Background()

	client := connectToK8s(*kubeconfig)
	deployClient := client.Resource(schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments", // 这里对应的是 api 请求中路径上的段名
	}).Namespace(metav1.NamespaceDefault)

	fmt.Println("Creating deployment...")
	if err := createDeploy(ctx, deployClient); err != nil {
		panic(err)
	}
	prompt()

	fmt.Println("Updating deployment...")
	if err := updateDeploy(ctx, deployClient); err != nil {
		panic(err)
	}
	prompt()

	fmt.Println("Listing deployment...")
	deployList, err := listDeploy(ctx, deployClient)
	if err != nil {
		panic(err)
	}
	for _, deploy := range deployList.Items {
		fmt.Println("deployList", deploy.GetName())
	}
	prompt()

	fmt.Println("Deleting deployment...")
	deleteDeploy(ctx, deployClient)
}

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
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

func prompt() {
	fmt.Println("> Press Enter to continue")
	fmt.Scanln()
}

func int32Ptr(v int32) *int32 {
	return &v
}

const (
	deployName = "foobar"
)

func createDeploy(ctx context.Context, deployClient typedappsv1.DeploymentInterface) error {
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deployName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "foobar",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "foobar",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:alpine",
						},
					},
				},
			},
		},
	}

	newDeploy, err := deployClient.Create(ctx, deploy, metav1.CreateOptions{})
	if err != nil {
		return errors.New("failed to create deploy")
	}

	fmt.Printf("%s\n", newDeploy.Name)

	return nil
}

func updateDeploy(ctx context.Context, deployClient typedappsv1.DeploymentInterface) error {
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		deploy, err := getDeploy(ctx, deployClient)
		if err != nil {
			return fmt.Errorf("updateDeploy fail: %w", err)
		}

		deploy.Spec.Replicas = int32Ptr(3)

		_, err = deployClient.Update(ctx, deploy, metav1.UpdateOptions{})
		return err
	})
	return err
}

func deleteDeploy(ctx context.Context, deployClient typedappsv1.DeploymentInterface) error {
	return deployClient.Delete(ctx, deployName, metav1.DeleteOptions{})
}

func listDeploy(
	ctx context.Context,
	deployClient typedappsv1.DeploymentInterface,
) (*appsv1.DeploymentList, error) {
	return deployClient.List(ctx, metav1.ListOptions{})
}

func getDeploy(
	ctx context.Context,
	deployClient typedappsv1.DeploymentInterface,
) (*appsv1.Deployment, error) {
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

	clientset := connectToK8s(*kubeconfig)
	deployClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)

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
		fmt.Println("deployList", deploy.Name)
	}
	prompt()

	fmt.Println("Deleting deployment...")
	deleteDeploy(ctx, deployClient)
}

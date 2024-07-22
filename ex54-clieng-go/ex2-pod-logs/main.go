package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
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

func getPodLogs(clientset *kubernetes.Clientset, podName, containerName string) {
	pods := clientset.CoreV1().Pods("default")
	opts := corev1.PodLogOptions{
		Container: containerName,
	}
	request := pods.GetLogs(podName, &opts)
	podLogs, err := request.Stream(context.TODO()) // 这里才开始实际执行
	if err != nil {
		log.Fatalln("error in opening stream.")
	}
	defer podLogs.Close()

	for {
		buf := make([]byte, 1024)
		num, err := podLogs.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("read err.")
		} else if num == 0 {
			continue
		} else {
			fmt.Println(string(buf[:num]))
		}
	}
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

	getPodLogs(clientset, "cppcheck-job-kjn7w", "")
}

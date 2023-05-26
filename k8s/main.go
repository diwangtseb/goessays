package main

import (
	"fmt"

	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 获取 kubeconfig 文件路径
	kubeconfig := "/path"

	// 使用 kubeconfig 创建一个 Config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	// 使用 Config 创建 Kubernetes 的客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 获取集群信息
	nsl, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, ns := range nsl.Items {
		fmt.Printf("Namespace name: %s\n", ns.ObjectMeta.Name)
	}
}

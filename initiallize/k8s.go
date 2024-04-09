package initiallize

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"mygin/client"
)

func InitK8SConnect() {
  var TLSConfig = rest.TLSClientConfig{
		Insecure: true,   // 忽略 CA 证书认证
	}

	config := &rest.Config{
		Host: "https://192.168.56.10:6443",
		BearerToken: "",
		TLSClientConfig: TLSConfig,
	}

	K8SClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	client.K8SClientSet = K8SClient

	nodeList, err := K8SClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, item := range nodeList.Items {
		fmt.Println(item.Namespace, item.Name)
	}


}
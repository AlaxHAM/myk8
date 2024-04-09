package client

import (
	"k8s.io/client-go/kubernetes"
)

var (
	K8SClientSet *kubernetes.Clientset
)
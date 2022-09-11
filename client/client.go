package client

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

// GetClient returns a Kubernetes clientset
func GetClient() *kubernetes.Clientset {

	config, err := genericclioptions.NewConfigFlags(true).ToRESTConfig()
	if err != nil {
		panic(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return client
}

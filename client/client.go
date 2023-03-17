package client

import (
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetLocal reads a user's KUBECONFIG file and returns a Client interface, a REST interface, and current namespace
func GetLocal() (*kubernetes.Clientset, *rest.Config, error) { //string, error) {
	configLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configPath := ""
	if configPath != "" {
		configPathList := filepath.SplitList(configPath)
		if len(configPathList) <= 1 {
			configLoadingRules.ExplicitPath = "~/.kube/config"
		} else {
			configLoadingRules.Precedence = configPathList
		}
	}

	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		configLoadingRules,
		&clientcmd.ConfigOverrides{
			CurrentContext: "",
		},
	)

	newConfig, err := config.ClientConfig()
	if err != nil {

		panic(err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(newConfig)
	if err != nil {
		log.Fatal(err)
	}

	return clientset, newConfig, nil
}

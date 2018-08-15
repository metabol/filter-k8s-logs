package main

import (
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func filter(line, kubeconfig, namespace string) string {
	filtered := ""

	client, err := getKubeClient(kubeconfig)
	if err != nil {
		log.Fatalf("cannot get Kubernetes client: %v", err)
	}

	secrets, err := getKubeSecrets(client)
	if err != nil {
		log.Fatalf("cannot get Kubernetes secrets: %v", err)
	}

	return filterSecrets(line, secrets)
}

func filterSecrets(line string, secrets v1.SecretList) string {

	return ""
}

func getKubeClient(kubeConfigLocation string) (*kubernetes.Clientset, error) {
	// build the config from the master and kubeconfig location
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigLocation)
	if err != nil {
		return nil, err
	}

	// creates the clientset
	return kubernetes.NewForConfig(config)
}

func getKubeSecrets(client *kubernetes.Clientset) (v1.SecretList, error) {
	return v1.SecretList{}, nil
}

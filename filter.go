package main

import (
	"log"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func filter(line, kubeconfig, namespace string) string {
	client, err := getKubeClient(kubeconfig)
	if err != nil {
		log.Fatalf("cannot get Kubernetes client: %v", err)
	}

	secrets, err := getKubeSecrets(client, namespace)
	if err != nil {
		log.Fatalf("cannot get Kubernetes secrets: %v", err)
	}

	for _, s := range secrets.Items {

		// ignore the default token secret for now
		//
		// TODO - add flag whether to include or exclude it
		if strings.Contains(s.Name, "default-token") {
			continue
		}

		for _, v := range s.Data {

			//fmt.Printf("log: %v; secret to check against: %s\n\n", line, string(v))

			// if the log line contains a secret value trim it
			if strings.Contains(line, string(v)) {
				line = strings.Replace(line, string(v), "[ redacted ]", -1)
			}
		}
	}

	return line
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

func getKubeSecrets(client *kubernetes.Clientset, namespace string) (*v1.SecretList, error) {
	opts := metav1.ListOptions{}
	return client.CoreV1().Secrets(namespace).List(opts)
}

package main

import (
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func filter(line string, secrets []v1.Secret) string {
	for _, s := range secrets {

		// ignore the default token secret for now
		//
		// TODO - add flag whether to include or exclude it
		if strings.Contains(s.Name, "default-token") {
			continue
		}

		for _, v := range s.Data {
			// if the log line contains a secret value redact it
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

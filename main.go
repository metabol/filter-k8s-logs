package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
	"github.com/radu-matei/filter-kubernetes-logs/cache"
)

func main() {
	logFile := getEnvVarOrExit("LOGS_FILE")
	namespace := getEnvVarOrExit("NAMESPACE")
	kubeconfig := getEnvVarOrExit("KUBECONFIG")

	client, err := getKubeClient(kubeconfig)
	if err != nil {
		log.Fatalf("cannot get Kubernetes client: %v", err)
	}

	cache := cache.New(client, namespace, cache.DefaultCacheSyncTimeout)
	secrets, err := cache.ListSecrets()
	if err != nil {
		log.Fatalf("cannot get Kubernetes secrets: %v", err)
	}

	t, err := tail.TailFile(logFile, tail.Config{Follow: true})
	if err != nil {
		log.Fatalf("cannot tail file: %v", err)
	}

	// this checks for every line in the log file
	//
	// TODO - implement checking in chunks of multiple lines
	for line := range t.Lines {
		fmt.Println(filter(line.Text, secrets))
	}
}

func getEnvVarOrExit(env string) string {
	val := os.Getenv(env)
	if val == "" {
		fmt.Printf("Missing environment variable %s\n", env)
		//os.Exit(1)
	}
	return val
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
)

var (
	logFile    string
	namespace  string
	kubeconfig string
)

func main() {
	logFile = getEnvVarOrExit("LOG_FILE")
	namespace = getEnvVarOrExit("NAMESPACE")
	kubeconfig = getEnvVarOrExit("KUBECONFIG")

	t, err := tail.TailFile(logFile, tail.Config{Follow: true})
	if err != nil {
		log.Fatalf("cannot tail file: %v", err)
	}

	for line := range t.Lines {
		fmt.Printf("filtered: %v", filter(line.Text, kubeconfig, namespace))
	}
}

func getEnvVarOrExit(env string) string {
	val := os.Getenv(env)
	if val == "" {
		fmt.Printf("Missing environment variable %s\n", env)
		os.Exit(1)
	}
	return val
}

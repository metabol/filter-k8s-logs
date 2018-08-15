package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
)

func main() {
	logFile := getEnvVarOrExit("LOG_FILE")
	namespace := getEnvVarOrExit("NAMESPACE")
	kubeconfig := getEnvVarOrExit("KUBECONFIG")

	t, err := tail.TailFile(logFile, tail.Config{Follow: false})
	if err != nil {
		log.Fatalf("cannot tail file: %v", err)
	}

	// this checks for every line in the log file
	//
	// TODO - implement checking in chunks of multiple lines
	for line := range t.Lines {
		fmt.Println(filter(line.Text, kubeconfig, namespace))
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

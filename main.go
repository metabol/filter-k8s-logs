package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	logFile := getEnvVarOrExit("LOG_FILE")

	f, err := os.Open(logFile)
	if err != nil {
		fmt.Printf("cannot open log file %v: %v", logFile, err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		// TODO - actual filtering

		fmt.Printf("sidecar log: %v", s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
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

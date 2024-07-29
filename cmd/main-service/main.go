package main

import (
	"flag"
	start "github.com/martirosharutyunyan/fraudnetic-test-task/internal"
	"log"
)

func main() {
	var env, host, port string

	flag.StringVar(&env, "env", "development", "Environment Variables filename")
	flag.StringVar(&host, "host", "localhost", "HTTP Server Address")
	flag.StringVar(&port, "port", "3000", "HTTP Server Address")
	flag.Parse()

	errC, err := start.Bootstrap(env, host, port)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

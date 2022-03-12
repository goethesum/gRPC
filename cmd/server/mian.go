package main

import "log"

func Run() error {
	// responsible for initilazing and starting
	// our gRPC server
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

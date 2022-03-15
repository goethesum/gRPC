package main

import (
	"log"

	"github.com/goethesum/gRPC/internal/db"
	"github.com/goethesum/gRPC/internal/rocket"
)

func Run() error {
	// responsible for initilazing and starting
	// our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

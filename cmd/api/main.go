package main

import (
	"log"

	"canipaddle.com/internal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("cannot initialize server: %s", err)
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("cannot initialize server: %s", err)
	}
}

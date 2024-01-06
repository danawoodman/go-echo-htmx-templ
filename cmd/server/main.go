package main

import (
	"log"

	"github.com/danawoodman/go-echo-htmx-templ/handlers"
)

func main() {
	server := handlers.NewServer()
	log.Fatal(server.Start())
}

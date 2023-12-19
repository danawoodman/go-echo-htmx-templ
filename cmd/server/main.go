package main

import (
	"log"
	"os"

	"github.com/danawoodman/go-echo-htmx-templ/controllers/views"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1:7788"
	}
	server := views.NewHTMLController(host)
	log.Fatal(server.Start())
}

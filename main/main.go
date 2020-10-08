package main

import (
	"github.com/gophercises/cyoa"
	"log"
	"net/http"
)

func main () {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server, err := cyoa.PrepareServer()
	http.ListenAndServe(":8080", server.HandleStart())
	return err
}

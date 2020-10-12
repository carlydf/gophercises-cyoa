package main

import (
	"flag"
	"github.com/gophercises/cyoa"
	"log"
	"net/http"
)

func main () {
	port := flag.String("port", "3000", "port to start CYOA app on")
	filename := flag.String("file", "../gopher.json", "JSON file with CYOA story")
	flag.Parse()
	run(filename, *port)
}

func run(filename *string, port string) error {
	h, err := cyoa.PrepareHandler(filename)
	if err != nil {
		return err
	}
	log.Fatal(http.ListenAndServe(":" + port, h))
	return nil
}

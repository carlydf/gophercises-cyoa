package main

import (
	"flag"
	"fmt"
	"github.com/gophercises/cyoa"
	"log"
	"net/http"
)

func main () {
	filename := flag.String("file", "../gopher.json", "JSON file with CYOA story")
	err := run(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func run(filename *string) error {
	server, err := cyoa.PrepareServer(filename)
	fmt.Printf("%+v", server.StoryArcs)
	http.ListenAndServe(":8080", server.HandleStart())
	return err
}

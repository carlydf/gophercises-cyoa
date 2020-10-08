package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type server struct {
	storyArcs map[string]StoryArc
	router *http.ServeMux
	// represents the service and holds all of its dependencies
	// the fields of the struct are shared dependencies
}

type StoryArc struct {
	Title string
	Story []string
	Options []map[string]string
}

func newServer() *server {
	s := &server{}
	s.router = http.NewServeMux()
	s.routes()
	return s
}

func PrepareServer() (*server, error) {
	jsonBytes, err := ioutil.ReadFile("../gopher.json")
	var arcs map[string]StoryArc
	err2 := json.Unmarshal(jsonBytes, &arcs)
	s := newServer()
	s.storyArcs = arcs
	if err2 != nil {
		return s, err2
	}
	return s, err
}



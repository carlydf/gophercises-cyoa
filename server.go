package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type server struct {
	StoryArcs map[string]StoryArc
	router *http.ServeMux
	// represents the service and holds all of its dependencies
	// the fields of the struct are shared dependencies
}

type StoryArc struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []map[string]string `json:"options"`
}

func newServer() *server {
	s := &server{}
	s.router = http.NewServeMux()
	s.routes()
	return s
}

func PrepareServer(filename *string) (*server, error) {
	jsonBytes, err := ioutil.ReadFile(*filename)
	var arcs map[string]StoryArc
	err2 := json.Unmarshal(jsonBytes, &arcs)
	s := newServer()
	s.StoryArcs = arcs
	if err2 != nil {
		return s, err2
	}
	return s, err
}



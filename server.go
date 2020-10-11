package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type server struct {
	Chapters map[string]Chapter
	router *http.ServeMux
	// represents the service and holds all of its dependencies
	// the fields of the struct are shared dependencies
}

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	NextChapter string `json:"arc"`
}

func newServer() *server {
	s := &server{}
	s.router = http.NewServeMux()
	s.routes()
	return s
}

func PrepareServer(filename *string) (*server, error) {
	jsonBytes, err := ioutil.ReadFile(*filename)
	var chapters map[string]Chapter
	err2 := json.Unmarshal(jsonBytes, &chapters)
	s := newServer()
	s.Chapters = chapters
	if err2 != nil {
		return s, err2
	}
	return s, err
}



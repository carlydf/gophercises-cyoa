package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
                <li><a href="/{{.NextChapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </body>
</html>
`

type server struct {
	Chapters map[string]Chapter
	router *http.ServeMux
	// represents the service and holds all of its dependencies
	// the fields of the struct are shared dependencies
}
func newServer() *server {
	s := &server{}
	s.router = http.NewServeMux()
	s.routes()
	return s
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

func PrepareServer(filename *string) (*server, error) {
	//htmlBytes, _ := ioutil.ReadFile("defaultHandlerTemplate.html")
	//defaultHandlerTmpl :=
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



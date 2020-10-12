package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
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
var tmpl *template.Template
//regardless of how many times that package is imported, the init() function will only be called once.
func init() {
	//Must says that if template doesnt compile correctly, there's no reason to continue
	tmpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

type Story map[string]Chapter

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	NextChapter string `json:"arc"`
}

type handler struct {
	s Story
}

func newHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	path := r.URL.Path
	chapRequested := "intro"
	if !(path == "" || path == "/") {
		chapRequested = r.URL.Path[1:] // take off the slash
	}
	if chapStruct, ok := h.s[chapRequested]; ok {
		err := tmpl.Execute(w, chapStruct)
		if err != nil {
			panic(err) //replace later? fine for dev
		}
	} else {
		fmt.Fprintf(w, "we could not find the chapter called %s", chapRequested)
	}
}

func PrepareHandler(filename *string) (http.Handler, error) {
	// htmlBytes, _ := ioutil.ReadFile("defaultHandlerTemplate.html")
	// defaultHandlerTmpl :=
	jsonBytes, err := ioutil.ReadFile(*filename)
	var s map[string]Chapter
	err2 := json.Unmarshal(jsonBytes, &s)
	h := newHandler(s)
	if err2 != nil {
		return h, err2
	}
	return h, err
}



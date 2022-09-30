package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func init() {
	tpl = template.Must(template.New("").Parse(html))
}

var tpl *template.Template

var html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Adventure</title>
</head>
<body>
    <h1>{{ .Title }}</h1>
	{{range .Paragraphs}}
		<p>{{.}}</p>
	{{end}}
	<ul>
		{{range .Options}}
			<li><a href="/{{$}}">{{.Text}}</li>
		{{end}}
	</ul>
</body>
</html>
`

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type handler struct {
	s Story
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)

			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	fmt.Printf("%v", h.s[path])
	fmt.Printf("%v", path)
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	var story Story
	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

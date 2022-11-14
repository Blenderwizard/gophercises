package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

var defautlTpml = `<!DOCTYPE html>
<html>
	<head>
		<title>
			Choose your own adventure
		</title>
	</head>
	<body>
		<h1>
			{{.Title}}
		</h1>
		{{range .Paragraphs}}
		<p>
			{{.}}
		</p>
		{{end}}
		<p>
			{{range .Options}}
			<p>
				<a href="/{{.Chapter}}">{{.Text}}</a>
			</p>
			{{end}}
		</p>
	</body>
</html>
`

func MapHandler(story Story, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("").Parse(defautlTpml))
		path := r.URL.Path[1:]
		if path == "" {
			path = "intro"
		}
		if chapter, ok := story[path]; ok {
			err := tmpl.Execute(w, chapter)
			if err != nil {
				http.Error(w, "Something went wrong...", http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "Chapter not found.", http.StatusNotFound)
	}
}

func JSONHandler(jsonFile io.Reader, fallback http.Handler) (http.HandlerFunc, error) {
	d := json.NewDecoder(jsonFile)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return MapHandler(story, fallback), nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

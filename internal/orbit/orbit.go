package orbit

import (
	"net/http"
	"path"

	"github.com/flosch/pongo2/v6"
)

// global state and http helper methods
type Orbit struct{}

func (o *Orbit) Text(w http.ResponseWriter, code int, text string) {
	w.Header().Add("content-type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(text))
}

func (o *Orbit) Html(w http.ResponseWriter, code int, html string) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(html))
}

func (o *Orbit) Render(w http.ResponseWriter, name string, code int, data pongo2.Context) {
	var template *pongo2.Template
	filename := path.Join("web/view", name)
	template = pongo2.Must(pongo2.FromCache(filename + "." + "ehtml"))

	w.Header().Add("content-type", "text/html")
	w.WriteHeader(code)
	template.ExecuteWriter(data, w)
}

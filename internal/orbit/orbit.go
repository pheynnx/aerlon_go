package orbit

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/ArminasAer/aerlon/internal/blogcache"
	"github.com/flosch/pongo2/v6"
)

// global state and http helper methods
type Orbit struct{}

func Launch(r http.Handler) {
	fmt.Printf("🚀 Aerlon launching: %s:%s 🚀\n", os.Getenv("HOST"), os.Getenv("PORT"))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), r)
	if err != nil {
		log.Fatal(err)
	}
}

func (o *Orbit) Text(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(text))
}

func (o *Orbit) HTML(w http.ResponseWriter, code int, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(html))
}

func (o *Orbit) Render(w http.ResponseWriter, name string, code int, data pongo2.Context) {
	template := pongo2.Must(pongo2.FromCache(path.Join("web/view", name) + ".ehtml"))

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	template.ExecuteWriter(data, w)
}

func (o *Orbit) CacheRender(w http.ResponseWriter, bc *blogcache.BlogCache, code int, slug string) {
	post, ok := bc.Posts[slug]
	if !ok {
		o.Error(w, 404, fmt.Sprintf("%s is not found", slug))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(post))
}

func (o *Orbit) Error(w http.ResponseWriter, code int, errorMessage string) {
	http.Error(w, errorMessage, code)
}

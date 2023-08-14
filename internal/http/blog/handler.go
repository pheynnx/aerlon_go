package blog

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BlogHandler struct {
	*BlogRouter
}

func (bh *BlogHandler) getBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		// bh.Orbit.Render(w, "blog", 200, map[string]any{"set": "hi"})

		if id == "42" {
			bh.blogCache.UpdateCache()
		}

		bh.Orbit.Html(w, 200, bh.blogCache.Index[2])
	}
}

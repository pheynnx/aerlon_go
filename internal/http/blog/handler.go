package blog

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BlogHandler struct {
	*BlogRouter
}

func (bh *BlogHandler) getBlogIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bh.Orbit.Html(w, 200, bh.blogCache.IndexMeta)
	}
}

func (bh *BlogHandler) getBlogBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		bh.Orbit.CacheRender(w, bh.blogCache, 200, slug)
	}
}

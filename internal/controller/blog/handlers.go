package blog

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*Router
}

func (h *Handler) getBlogIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.Orbit.HTML(w, 200, h.blogCache.IndexMeta)
	}
}

func (h *Handler) getBlogBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		h.Orbit.CacheRender(w, h.blogCache, 200, slug)
	}
}

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

		c, err := r.Cookie("layout")
		if err != nil {
			h.Orbit.HTML(w, http.StatusOK, h.blogCache.IndexMeta)
			return
		}

		if c.Value == "compact" {
			h.Orbit.HTML(w, http.StatusOK, h.blogCache.IndexCompactMeta)
			return
		}

		h.Orbit.HTML(w, http.StatusOK, h.blogCache.IndexMeta)
	}
}

func (h *Handler) getBlogBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		h.Orbit.CacheRender(w, h.blogCache, http.StatusOK, slug)
	}
}

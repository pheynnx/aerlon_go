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

		h.Orbit.HTML(w, http.StatusOK, h.htmlStore.StaticIndex)
	}
}

func (h *Handler) getBlogBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		h.Orbit.StaticRender(w, h.htmlStore, http.StatusOK, slug)
	}
}

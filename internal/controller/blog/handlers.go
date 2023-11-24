package blog

import (
	"fmt"
	"net/http"

	"github.com/ArminasAer/aerlon/internal/views"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*Router
}

func (h *Handler) getBlogIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c := views.IndexBuilder(h.postCache.TemplStore, true)
		c.Render(r.Context(), w)
	}
}

func (h *Handler) getBlogBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		post, ok := h.postCache.TemplPostMap[slug]
		if !ok {
			http.Error(w, fmt.Sprintf("%s is not found", slug), 404)
			return
		}

		c := views.BlogBuilder(post)
		c.Render(r.Context(), w)
	}
}

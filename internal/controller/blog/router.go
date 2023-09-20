package blog

import (
	"net/http"

	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/ArminasAer/aerlon/internal/static"
	"github.com/go-chi/chi/v5"
)

// blog router state
//
// extends chi router and orbit
type Router struct {
	htmlStore *static.HTMLStore
	*chi.Mux
	*orbit.Orbit
}

func newRouter(htmlStore *static.HTMLStore) *Router {
	return &Router{
		htmlStore: htmlStore, Mux: chi.NewRouter(),
	}
}

func Routes(htmlStore *static.HTMLStore) *Router {
	router := newRouter(htmlStore)

	h := Handler{router}

	router.Get("/", h.getBlogIndex())

	router.Route("/blog", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		})
		r.Get("/{slug}", h.getBlogBySlug())
	})

	return router
}

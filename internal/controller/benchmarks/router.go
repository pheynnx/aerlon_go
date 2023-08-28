package benchmarks

import (
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
	*orbit.Orbit
}

func newRouter() *Router {
	return &Router{
		Mux: chi.NewRouter(),
	}
}

func Routes() *Router {
	router := newRouter()

	h := Handler{router}

	router.Get("/", h.getBenchmarks())

	return router
}

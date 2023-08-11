package blog

import (
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

// blog router state
//
// extends chi router and orbit
type BlogRouter struct {
	// temp state
	DB string
	*chi.Mux
	*orbit.Orbit
}

func newBlogRouter(DB string) *BlogRouter {
	return &BlogRouter{
		DB: DB, Mux: chi.NewRouter(),
	}
}

func BlogRoutes(DB string) *BlogRouter {
	blogRouter := newBlogRouter(DB)

	bh := BlogHandler{BlogRouter: blogRouter}

	blogRouter.Get("/", bh.getBlog())

	return blogRouter
}

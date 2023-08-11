package blog

import (
	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

// blog router state
//
// extends chi router and orbit
type BlogRouter struct {
	// temp state
	DBs *database.PostgresPool
	*chi.Mux
	*orbit.Orbit
}

func newBlogRouter(DBs *database.PostgresPool) *BlogRouter {
	return &BlogRouter{
		DBs: DBs, Mux: chi.NewRouter(),
	}
}

func BlogRoutes(DBs *database.PostgresPool) *BlogRouter {
	blogRouter := newBlogRouter(DBs)

	bh := BlogHandler{BlogRouter: blogRouter}

	blogRouter.Get("/", bh.getBlog())

	return blogRouter
}

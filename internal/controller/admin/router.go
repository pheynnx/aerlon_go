package admin

import (
	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
	*orbit.Orbit
	DB *database.DBPool
}

func newRouter(DB *database.DBPool) *Router {
	return &Router{
		Mux: chi.NewRouter(),
		DB:  DB,
	}
}

func Routes(DB *database.DBPool) *Router {
	router := newRouter(DB)

	ah := Handler{router}

	router.Get("/login", ah.AdminLoginView())
	router.Get("/", ah.AdminDashboardView())

	return router
}

package admin

import (
	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

type AdminRouter struct {
	*chi.Mux
	*orbit.Orbit
	DB *database.DBPool
}

func newAdminRouter(DB *database.DBPool) *AdminRouter {
	return &AdminRouter{
		Mux: chi.NewRouter(),
		DB:  DB,
	}
}

func AdminRoutes(DB *database.DBPool) *AdminRouter {
	AdminRouter := newAdminRouter(DB)

	ah := AdminHandler{AdminRouter: AdminRouter}

	AdminRouter.Get("/login", ah.AdminLoginView())
	AdminRouter.Get("/", ah.AdminDashboardView())

	return AdminRouter
}

package admin

import (
	"net/http"

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

	h := Handler{router}

	router.Group(func(r chi.Router) {
		r.Use(AdminLoginCheckAuth())
		r.Get("/", h.AdminDashboardView())
	})

	router.Group(func(r chi.Router) {
		r.Use(AdminIfAuthenticated())
		r.Get("/login", h.AdminLoginView())
	})

	router.Route("/api", func(r chi.Router) {
		r.Post("/login", h.AdminAPILogin())
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {})
		r.Patch("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {})
		r.Delete("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {})
	})

	return router
}

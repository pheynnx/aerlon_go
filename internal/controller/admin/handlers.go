package admin

import "net/http"

type Handler struct {
	*Router
}

func (h Handler) AdminLoginView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.Orbit.Render(w, "admin/login", 200, map[string]any{})
	}
}

func (h Handler) AdminDashboardView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.Orbit.Render(w, "admin/dashboard", 200, map[string]any{})
	}
}

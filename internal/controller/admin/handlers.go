package admin

import "net/http"

type AdminHandler struct {
	*AdminRouter
}

func (ah AdminHandler) AdminLoginView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ah.Orbit.Render(w, "admin/login", 200, map[string]any{})
	}
}

func (ah AdminHandler) AdminDashboardView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ah.Orbit.Render(w, "admin/dashboard", 200, map[string]any{})
	}
}

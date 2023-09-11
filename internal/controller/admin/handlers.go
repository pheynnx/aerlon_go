package admin

import (
	"net/http"
	"os"
)

type Handler struct {
	*Router
}

// admin

func (h Handler) AdminLoginView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.Orbit.Render(w, "admin/login", http.StatusOK, map[string]any{})
	}
}

func (h Handler) AdminDashboardView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.Orbit.Render(w, "admin/dashboard", http.StatusOK, map[string]any{})
	}
}

// /admin/api
func (h Handler) AdminAPILogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("password")
		pin := r.FormValue("pin")

		if password == os.Getenv("ADMIN_PASSWORD") && pin == os.Getenv("ADMIN_PIN") {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}

func (h Handler) AdminAPILogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h Handler) AdminAPIGetPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h Handler) AdminAPIGetPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h Handler) AdminAPICreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h Handler) AdminAPIUpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h Handler) AdminAPIDeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

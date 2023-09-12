package admin

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ArminasAer/aerlon/internal/auth"
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

			token, err := auth.CreateAdminJWT(password, pin)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusForbidden)
				return
			}

			cookie := http.Cookie{
				Name:     "auth",
				Value:    token,
				Path:     "/admin",
				HttpOnly: true,
				Expires:  time.Now().Add(96 * time.Hour),
			}

			http.SetCookie(w, &cookie)

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

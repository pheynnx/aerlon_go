package admin

import (
	"net/http"

	"github.com/ArminasAer/aerlon/internal/auth"
)

func AdminLoginCheckAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			c, err := r.Cookie("auth")
			if err != nil {
				http.Redirect(w, r, "/admin/login", http.StatusFound)
				return
			}

			token := c.Value

			if auth.VerifyAdminJWT(token) {
				next.ServeHTTP(w, r)
				return
			}
		})
	}
}

func AdminIfAuthenticated() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			c, err := r.Cookie("auth")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			token := c.Value

			if auth.VerifyAdminJWT(token) {
				http.Redirect(w, r, "/admin", http.StatusFound)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func AdminAPICheckAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CHECK FOR AUTH

			// IF AUTH GO TO /ADMIN

			// IF NO AUTH GO TO /LOGIN
			// w.WriteHeader(http.StatusForbidden)
			// return

			next.ServeHTTP(w, r)
		})
	}
}

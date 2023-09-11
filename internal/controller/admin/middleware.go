package admin

import "net/http"

func AdminLoginCheckAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CHECK FOR AUTH

			// IF AUTH GO TO /ADMIN
			// next.ServeHTTP(w, r)

			// IF NO AUTH GO TO /LOGIN
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		})
	}
}

func AdminIfAuthenticated() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// IF AUTH TRUE
			// http.Redirect(w, r, "/admin", http.StatusFound)

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

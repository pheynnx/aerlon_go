package blog

import (
	"net/http"

	"github.com/ArminasAer/aerlon/internal/blogcache"
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

// blog router state
//
// extends chi router and orbit
type Router struct {
	blogCache *blogcache.BlogCache
	*chi.Mux
	*orbit.Orbit
}

func newRouter(blogCache *blogcache.BlogCache) *Router {
	return &Router{
		blogCache: blogCache, Mux: chi.NewRouter(),
	}
}

func Routes(blogCache *blogcache.BlogCache) *Router {
	router := newRouter(blogCache)

	h := Handler{router}

	router.Get("/", h.getBlogIndex())
	router.Get("/switch", func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("layout")
		if err != nil {

			cookie := http.Cookie{
				Name:     "layout",
				Value:    "compact",
				Path:     "/",
				HttpOnly: false,
			}

			http.SetCookie(w, &cookie)

			h.Orbit.Render(w, "partials/post_container_compact", http.StatusOK, map[string]any{"metaList": h.blogCache.MetaCache})
			return
		}

		if c.Value == "compact" {

			cookie := http.Cookie{
				Name:     "layout",
				Value:    "full",
				Path:     "/",
				HttpOnly: false,
			}

			http.SetCookie(w, &cookie)

			h.Orbit.Render(w, "partials/post_container", http.StatusOK, map[string]any{"metaList": h.blogCache.MetaCache})
			return
		}

		cookie := http.Cookie{
			Name:     "layout",
			Value:    "compact",
			Path:     "/",
			HttpOnly: false,
		}

		http.SetCookie(w, &cookie)

		h.Orbit.Render(w, "partials/post_container_compact", http.StatusOK, map[string]any{"metaList": h.blogCache.MetaCache})
	})

	router.Route("/blog", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		})
		r.Get("/{slug}", h.getBlogBySlug())
	})

	// middleware testing
	// router.Route("/weha", func(r chi.Router) {
	// 	r.Use(func(next http.Handler) http.Handler {
	// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 			next.ServeHTTP(w, r)
	// 		})
	// 	})
	// 	r.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("HEYA"))
	// 	})
	// })

	return router
}

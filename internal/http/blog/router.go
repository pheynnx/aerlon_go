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
type BlogRouter struct {
	blogCache *blogcache.BlogCache
	*chi.Mux
	*orbit.Orbit
}

func newBlogRouter(blogCache *blogcache.BlogCache) *BlogRouter {
	return &BlogRouter{
		blogCache: blogCache, Mux: chi.NewRouter(),
	}
}

func BlogRoutes(blogCache *blogcache.BlogCache) *BlogRouter {
	blogRouter := newBlogRouter(blogCache)

	bh := BlogHandler{BlogRouter: blogRouter}

	// blogRouter.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		next.ServeHTTP(w, r)
	// 	})
	// })

	blogRouter.Get("/", bh.getBlogIndex())

	blogRouter.Route("/blog", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		})
		r.Get("/{slug}", bh.getBlogBySlug())
	})

	// middleware testing
	blogRouter.Route("/weha", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		})
		r.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("HEYA"))
		})
	})

	return blogRouter
}

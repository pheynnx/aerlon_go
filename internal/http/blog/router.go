package blog

import (
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

	blogRouter.Get("/", bh.getBlogIndex())
	blogRouter.Get("/{slug}", bh.getBlogBySlug())

	return blogRouter
}

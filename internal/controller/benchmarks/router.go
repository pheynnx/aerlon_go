package benchmarks

import (
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

type BenchmarksRouter struct {
	*chi.Mux
	*orbit.Orbit
}

func newBenchmarksRouter() *BenchmarksRouter {
	return &BenchmarksRouter{
		Mux: chi.NewRouter(),
	}
}

func BenchmarksRoutes() *BenchmarksRouter {
	BenchmarksRouter := newBenchmarksRouter()

	bmh := BenchmarksHandler{BenchmarksRouter: BenchmarksRouter}

	BenchmarksRouter.Get("/", bmh.getBenchmarks())

	return BenchmarksRouter
}

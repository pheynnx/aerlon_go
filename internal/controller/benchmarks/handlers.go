package benchmarks

import "net/http"

type BenchmarksHandler struct {
	*BenchmarksRouter
}

func (bmh BenchmarksHandler) getBenchmarks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bmh.Orbit.Render(w, "benchmarks", 200, map[string]any{})
	}
}

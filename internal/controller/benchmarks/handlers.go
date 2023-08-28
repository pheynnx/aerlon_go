package benchmarks

import "net/http"

type Handler struct {
	*Router
}

func (h Handler) getBenchmarks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.Orbit.Render(w, "benchmarks", 200, map[string]any{})
	}
}

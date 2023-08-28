package readme

import "net/http"

type Handler struct {
	*Router
}

func (h Handler) getReadme() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.Orbit.Render(w, "readme", 200, map[string]any{})
	}
}

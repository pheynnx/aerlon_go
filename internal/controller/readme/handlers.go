package readme

import "net/http"

type ReadmeHandler struct {
	*ReadmeRouter
}

func (rh ReadmeHandler) getReadme() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rh.Orbit.Render(w, "readme", 200, map[string]any{})
	}
}

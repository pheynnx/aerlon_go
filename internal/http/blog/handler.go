package blog

import (
	"net/http"
)

type BlogHandler struct {
	*BlogRouter
}

func (bh *BlogHandler) getBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bh.Orbit.Render(w, "blog", 200, map[string]any{"set": "hi"})
	}
}

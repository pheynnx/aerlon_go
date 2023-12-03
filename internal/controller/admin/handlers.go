package admin

import (
	"net/http"

	"github.com/ArminasAer/aerlon/internal/views/admin"
)

type Handler struct {
	*Router
}

func (h *Handler) getAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := admin.AdminLoginBuilder()
		c.Render(r.Context(), w)
	}
}

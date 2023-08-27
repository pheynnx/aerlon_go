package readme

import (
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

type ReadmeRouter struct {
	*chi.Mux
	*orbit.Orbit
}

func newReadmeRouter() *ReadmeRouter {
	return &ReadmeRouter{
		Mux: chi.NewRouter(),
	}
}

func ReadMeRoutes() *ReadmeRouter {
	readmeRouter := newReadmeRouter()

	rh := ReadmeHandler{ReadmeRouter: readmeRouter}

	readmeRouter.Get("/", rh.getReadme())

	return readmeRouter
}

package station

import (
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

type StationRouter struct {
	*chi.Mux
	*orbit.Orbit
}

func newStationRouter() *StationRouter {
	return &StationRouter{
		Mux: chi.NewRouter(),
	}
}

func StationRoutes() *StationRouter {
	stationRouter := newStationRouter()

	sh := StationHandler{StationRouter: stationRouter}

	stationRouter.Get("/", sh.getStation())

	return stationRouter
}

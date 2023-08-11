package station

import "net/http"

type StationHandler struct {
	*StationRouter
}

func (sh *StationHandler) getStation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sh.Orbit.Render(w, "station", 200, map[string]any{})
	}
}

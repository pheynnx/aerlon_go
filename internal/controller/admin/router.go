package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	// postCache *cache.PostCache
	db *database.DBPool
	*chi.Mux
	*orbit.Orbit
}

func newRouter(db *database.DBPool) *Router {
	return &Router{
		db: db, Mux: chi.NewRouter(),
	}
}

func Routes(db *database.DBPool) *Router {
	router := newRouter(db)

	h := Handler{router}

	router.Get("/login", h.getAdmin())

	type Login struct {
		Password string `json:"password"`
		Pin      string `json:"pin"`
	}

	router.Post("/login/user", func(w http.ResponseWriter, r *http.Request) {
		login := Login{}

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&login)
		if err != nil {
			panic(err)
		}
		fmt.Println(login)
	})

	return router
}

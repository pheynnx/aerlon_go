package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/ArminasAer/aerlon/internal/controller/blog"
	"github.com/ArminasAer/aerlon/internal/orbit"
	"github.com/ArminasAer/aerlon/internal/static"
	"github.com/ArminasAer/aerlon/internal/views"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// initialize sqlx pool
	// db, err := database.NewDBPool()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// initialize blog html cache
	hs, err := static.InitStore()
	if err != nil {
		log.Fatal(err)
	}

	// create top level router
	r := chi.NewRouter()

	// root level middleware stack
	r.Use(chiMiddleware.RealIP)
	// r.Use(middleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// server static files
	// r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "web/root/favicon.ico")
	// })
	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/root/robots.txt")
	})
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// app routers

	r.Group(func(r chi.Router) {

		// r.Use(middleware.Metrics(db))

		r.Mount("/", blog.Routes(hs))
	})

	r.Get("/templ/test", func(w http.ResponseWriter, r *http.Request) {
		c := views.IndexBuilder(hs.TemplStore, true)
		c.Render(r.Context(), w)
	})

	// start server
	orbit.Launch(r)
}

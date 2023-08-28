package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/ArminasAer/aerlon/internal/blogcache"
	"github.com/ArminasAer/aerlon/internal/controller/admin"
	"github.com/ArminasAer/aerlon/internal/controller/benchmarks"
	"github.com/ArminasAer/aerlon/internal/controller/blog"
	"github.com/ArminasAer/aerlon/internal/controller/readme"
	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/orbit"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// initalize sqlx pool
	db, err := database.NewDBPool()
	if err != nil {
		log.Fatal(err)
	}

	// initalize blog html cache
	bc, err := blogcache.InitCache(db)
	if err != nil {
		log.Fatal(err)
	}

	// create top level router
	r := chi.NewRouter()

	// root level middleware stack
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// server static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// app routers
	r.Mount("/", blog.BlogRoutes(bc))
	r.Mount("/readme", readme.ReadMeRoutes())
	r.Mount("/benchmarks", benchmarks.BenchmarksRoutes())

	// admin router
	r.Mount("/admin", admin.AdminRoutes(db))

	// start server
	orbit.Launch(r)
}

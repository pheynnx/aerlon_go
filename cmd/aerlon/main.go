package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/ArminasAer/aerlon/internal/blogcache"
	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/http/blog"
	"github.com/ArminasAer/aerlon/internal/http/station"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// compile scss with 'sass'
	cmd := exec.Command("sass", "--no-source-map", "--style=compressed", "web/source/scss:web/static/css")
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
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

	// server static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// mount routers
	r.Mount("/", station.StationRoutes())
	r.Mount("/blog", blog.BlogRoutes(bc))

	// start server
	fmt.Printf("🚀 Aerlon launching: %s:%s 🚀\n", os.Getenv("HOST"), os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), r)
}

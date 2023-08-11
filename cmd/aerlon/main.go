package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/ArminasAer/aerlon/internal/http/blog"
	"github.com/ArminasAer/aerlon/internal/http/station"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
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

	// initalize database
	db := "Some DB"

	r := chi.NewRouter()

	// server static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// mount routers
	r.Mount("/", station.StationRoutes())
	r.Mount("/blog", blog.BlogRoutes(db))

	// start server
	fmt.Printf("🚀 Aerlon launching: %s:%s 🚀\n", os.Getenv("HOST"), os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), r)
}

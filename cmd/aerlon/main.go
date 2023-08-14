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

	// SQLXPool, err := database.NewSQLXPool()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var posts []model.Post
	// rows, err := SQLXPool.Queryx("SELECT * FROM post")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// for rows.Next() {
	// 	var p model.Post
	// 	err = rows.StructScan(&p)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	posts = append(posts, p)
	// }
	// for _, v := range posts {
	// 	fmt.Println(v.Categories)
	// }

	// var posts []model.Post
	// err = SQLXPool.Select(&posts, "SELECT * FROM post")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, v := range posts {
	// 	fmt.Println(v.Categories)
	// }

	bc := blogcache.InitCache()

	r := chi.NewRouter()

	// server static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// mount routers
	r.Mount("/", station.StationRoutes())
	r.Mount("/blog/{id}", blog.BlogRoutes(bc))

	// start server
	fmt.Printf("🚀 Aerlon launching: %s:%s 🚀\n", os.Getenv("HOST"), os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), r)
}

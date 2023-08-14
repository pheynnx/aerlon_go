package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"

	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/http/station"
)

type Post struct {
	Id            uuid.UUID
	Date          time.Time
	Slug          string
	Title         string
	Series        string
	Categories    pq.StringArray
	Markdown      string
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
	Published     bool
	Featured      bool
	PostSnippet   string `db:"post_snippet"`
	SeriesSnippet string `db:"series_snippet"`
}

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

	SQLXPool, err := database.NewSQLXPool()
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	rows, err := SQLXPool.Queryx("SELECT * FROM post")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var p Post
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Println(err.Error())
		}
		posts = append(posts, p)
	}
	for _, v := range posts {
		fmt.Println(v.Categories)
	}

	// var posts []Post
	// err = SQLXPool.Select(&posts, "SELECT * FROM post")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(posts)

	// rows, err := postPool.Query(context.Background(), "SELECT * FROM post")

	r := chi.NewRouter()

	// server static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// mount routers
	r.Mount("/", station.StationRoutes())
	// r.Mount("/blog", blog.BlogRoutes())

	// start server
	fmt.Printf("🚀 Aerlon launching: %s:%s 🚀\n", os.Getenv("HOST"), os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), r)
}

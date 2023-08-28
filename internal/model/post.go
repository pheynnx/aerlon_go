package model

import (
	"bytes"
	"sort"
	"time"

	"github.com/ArminasAer/aerlon/internal/database"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type Post struct {
	ID            uuid.UUID      `json:"id"`
	Date          time.Time      `json:"date"`
	Slug          string         `json:"slug"`
	Title         string         `json:"title"`
	Series        string         `json:"series"`
	Categories    pq.StringArray `json:"categories"`
	Markdown      string         `json:"markdown"`
	CreatedAt     string         `db:"created_at" json:"created_at"`
	UpdatedAt     string         `db:"updated_at" json:"updated_at"`
	Published     bool           `json:"published"`
	Featured      bool           `json:"featured"`
	PostSnippet   string         `db:"post_snippet" json:"post_snippet"`
	SeriesSnippet string         `db:"series_snippet" json:"series_snippet"`
}

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithFormatOptions(
				chromahtml.WithClasses(true),
			),
		),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
)

func (p *Post) ConvertMarkdownToHTML() error {
	var buf bytes.Buffer
	err := md.Convert([]byte(p.Markdown), &buf)
	if err != nil {
		return err
	}

	p.Markdown = buf.String()

	return nil
}

func SortPostsByDate(posts []*Post) {
	sort.Slice(posts, func(i, j int) bool {
		a := posts[i].Date.Unix()
		b := posts[j].Date.Unix()

		if a == b {
			return posts[i].Title < posts[j].Title
		} else if a > b {
			return a > b
		}
		return b < a
	})
}

func GetPostFromDB(DB *database.DBPool, id uuid.UUID) (*Post, error) {
	var post *Post
	err := DB.Get(&post, "SELECT * FROM post WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func GetPostsFromDB(DB *database.DBPool) ([]*Post, error) {
	var posts []*Post
	err := DB.Select(&posts, "SELECT * FROM post")
	if err != nil {
		return nil, err
	}

	return posts, nil
}

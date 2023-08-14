package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Post struct {
	Id            uuid.UUID
	Date          time.Time
	Slug          string
	Title         string
	Series        string
	Categories    pq.StringArray
	Markdown      string
	CreatedAt     string `db:"created_at" json:"created_at"`
	UpdatedAt     string `db:"updated_at" json:"updated_at"`
	Published     bool
	Featured      bool
	PostSnippet   string `db:"post_snippet" json:"post_snippet"`
	SeriesSnippet string `db:"series_snippet" json:"series_snippet"`
}

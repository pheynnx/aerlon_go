package model

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
	"time"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

type Post struct {
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Slug        string    `json:"slug"`
	Series      string    `json:"series"`
	Categories  []string  `json:"categories"`
	Markdown    string    `json:"markdown"`
	Published   bool      `json:"published"`
	Featured    bool      `json:"featured"`
	PostSnippet string    `db:"post_snippet" json:"post_snippet"`
}

var md = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithWrapperRenderer(func(w util.BufWriter, context highlighting.CodeBlockContext, entering bool) {
				lang, _ := context.Language()

				if entering {
					if lang == nil {
						w.WriteString("<pre><code>")
						return
					}
					w.WriteString(fmt.Sprintf(`<div class="code-block"><p class="code-block-header"><span class="language-name">%s</span></p><pre class="chroma"><code class="language-`, lang))
					w.Write(lang)
					w.WriteString(`" data-lang="`)
					w.Write(lang)
					w.WriteString(`">`)
				} else {
					if lang == nil {
						w.WriteString("</pre></code>")
						return
					}
					w.WriteString(`</code></pre></div>`)
				}
			}),
			highlighting.WithFormatOptions(
				chromahtml.PreventSurroundingPre(true),
				chromahtml.WithClasses(true),
			),
		),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
)

func NewPostArray() ([]*Post, error) {
	files, err := os.ReadDir("./posts")
	if err != nil {
		return nil, err
	}

	var posts []*Post

	for _, f := range files {
		content, err := os.ReadFile(fmt.Sprintf("./posts/%s", f.Name()))
		if err != nil {
			return nil, err
		}

		post, err := ParseMarkdownAndMeta(content)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	// Sort
	SortPostsByDate(posts)

	return posts, nil
}

func ParseMarkdownAndMeta(content []byte) (*Post, error) {
	var buf bytes.Buffer
	cxt := parser.NewContext()
	err := md.Convert(content, &buf, parser.WithContext(cxt))
	if err != nil {
		return nil, err
	}

	meta := meta.Get(cxt)

	date, err := time.Parse("January 2, 2006", meta["Date"].(string))
	if err != nil {
		return nil, err
	}

	var categories []string
	for _, c := range meta["Categories"].([]interface{}) {
		categories = append(categories, c.(string))
	}

	slices.Sort(categories)

	return &Post{
		Title:       meta["Title"].(string),
		Date:        date,
		Slug:        meta["Slug"].(string),
		Series:      meta["Series"].(string),
		Categories:  categories,
		Markdown:    buf.String(),
		Published:   meta["Published"].(bool),
		Featured:    meta["Featured"].(bool),
		PostSnippet: meta["PostSnippet"].(string),
	}, nil
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

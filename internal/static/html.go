package static

import (
	"path"

	"github.com/ArminasAer/aerlon/internal/model"
	"github.com/flosch/pongo2/v6"
)

type HTMLStore struct {
	StaticIndex   string
	StaticPostMap map[string]string
}

func InitStore() (*HTMLStore, error) {
	posts, err := model.NewPostArray()
	if err != nil {
		return nil, err
	}

	postList := []*model.Post{}

	postsRendered := map[string]string{}
	postsRenderer := pongo2.Must(pongo2.FromCache(path.Join("web/view", "blog_$post.ehtml")))

	for _, p := range posts {

		if p.Published {
			postList = append(postList, p)
		}

		pr, err := postsRenderer.Execute(map[string]any{"post": p})
		if err != nil {
			return nil, err
		}
		postsRendered[p.Slug] = pr
	}

	indexRenderer := pongo2.Must(pongo2.FromCache(path.Join("web/view", "index.ehtml")))
	blogIndexRendered, err := indexRenderer.Execute(map[string]any{"postList": postList})
	if err != nil {
		return nil, err
	}

	return &HTMLStore{
		StaticIndex:   blogIndexRendered,
		StaticPostMap: postsRendered,
	}, nil
}

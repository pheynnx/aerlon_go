package blogcache

import (
	"path"
	"slices"

	"github.com/ArminasAer/aerlon/internal/model"
	"github.com/flosch/pongo2/v6"
)

type BlogCache struct {
	IndexMeta        string
	IndexCompactMeta string
	Posts            map[string]string
	MetaCache        []*model.Post
}

func InitCache() (*BlogCache, error) {
	posts, err := model.NewPostArray()
	if err != nil {
		return nil, err
	}

	metaList := []*model.Post{}

	postsRendered := map[string]string{}
	postsRenderer := pongo2.Must(pongo2.FromCache(path.Join("web/view", "blog_$post.ehtml")))

	for _, p := range posts {

		// sort categories by alphabetical order
		slices.Sort(p.Categories)

		// split slice into two slices by featured bool

		if p.Published {
			metaList = append(metaList, p)

		}

		pr, err := postsRenderer.Execute(map[string]any{"postStruct": p})
		if err != nil {
			return nil, err
		}
		postsRendered[p.Slug] = pr

	}

	indexRenderer := pongo2.Must(pongo2.FromCache(path.Join("web/view", "index.ehtml")))
	blogIndexRendered, err := indexRenderer.Execute(map[string]any{"metaList": metaList, "url": "/"})
	if err != nil {
		return nil, err
	}

	indexCompactRenderer := pongo2.Must(pongo2.FromCache(path.Join("web/view", "index_compact.ehtml")))
	blogIndexCompactRendered, err := indexCompactRenderer.Execute(map[string]any{"metaList": metaList, "url": "/"})
	if err != nil {
		return nil, err
	}

	return &BlogCache{
		IndexMeta:        blogIndexRendered,
		IndexCompactMeta: blogIndexCompactRendered,
		Posts:            postsRendered,
		MetaCache:        metaList,
	}, nil
}

// SetInitCache on startup
// Needs to grab posts from postgres
// Needs to sort the posts by date
// Create a sort index (meta)
// Convert markdown to html
// Create a sorted (post)
// Cache needs to stay sorted by date

// func (bc *BlogCache) UpdateCache() {
// 	bc.mx.Lock()
// 	defer bc.mx.Unlock()

// 	// bc.Index[2] = "taco"
// 	// bc.Posts["keyOne"] = model.Post{}
// }

// Delete spot in Cache
// Updated Index and Posts

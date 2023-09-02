package blogcache

import (
	"path"
	"slices"
	"sync"

	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/model"
	"github.com/ArminasAer/aerlon/internal/model/dto"
	"github.com/flosch/pongo2/v6"
)

type BlogCache struct {
	mx        sync.Mutex
	IndexMeta string
	Posts     map[string]string
}

func InitCache(DB *database.DBPool) (*BlogCache, error) {

	// grab posts from DB
	posts, err := model.GetPostsFromDB(DB)
	if err != nil {
		return nil, err
	}

	// sort posts
	model.SortPostsByDate(posts)

	featuredMeta := []*dto.Meta{}
	nonFeaturedMeta := []*dto.Meta{}

	postsRendered := map[string]string{}
	postsRenderer := pongo2.Must(pongo2.FromCache(path.Join("web/view", "blog_$post.ehtml")))

	for _, p := range posts {

		err := p.ConvertMarkdownToHTML()
		if err != nil {
			return nil, err
		}

		// sort categories by alphabetical order
		slices.Sort(p.Categories)

		// split slice into two slices by featured bool
		if p.Featured {
			featuredMeta = append(featuredMeta, dto.MetaFromPost(p))
		} else {
			nonFeaturedMeta = append(nonFeaturedMeta, dto.MetaFromPost(p))
		}

		pr, err := postsRenderer.Execute(map[string]any{"postStruct": p})
		if err != nil {
			return nil, err
		}
		postsRendered[p.Slug] = pr

	}

	indexRenderer := pongo2.Must(pongo2.FromCache(path.Join("web/view", "index.ehtml")))
	blogIndexRendered, err := indexRenderer.Execute(map[string]any{"featuredMetaList": featuredMeta, "nonFeaturedMetaList": nonFeaturedMeta, "url": "/"})
	if err != nil {
		return nil, err
	}

	return &BlogCache{
		IndexMeta: blogIndexRendered,
		Posts:     postsRendered,
	}, nil
}

// SetInitCache on startup
// Needs to grab posts from postgres
// Needs to sort the posts by date
// Create a sort index (meta)
// Convert markdown to html
// Create a sorted (post)
// Cache needs to stay sorted by date

func (bc *BlogCache) UpdateCache() {
	bc.mx.Lock()
	defer bc.mx.Unlock()

	// bc.Index[2] = "taco"
	// bc.Posts["keyOne"] = model.Post{}
}

// Delete spot in Cache
// Updated Index and Posts

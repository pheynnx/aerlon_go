package blogcache

import (
	"sync"

	"github.com/ArminasAer/aerlon/internal/database"
	"github.com/ArminasAer/aerlon/internal/model"
)

type BlogCache struct {
	mx    sync.Mutex
	Index []string
	Posts map[string]model.Post
}

func InitCache(DB *database.DBPool) *BlogCache {

	model.GetPostsFromDB(DB)

	return &BlogCache{
		Index: []string{"one", "two", "three"},
		Posts: map[string]model.Post{},
	}
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

	bc.Index[2] = "taco"
	bc.Posts["keyOne"] = model.Post{}
}

// Delete spot in Cache
// Updated Index and Posts

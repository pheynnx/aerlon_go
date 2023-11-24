package cache

import (
	"github.com/ArminasAer/aerlon/internal/model"
)

// naming of the cache field will change
// also this cache store might be moved under the model package into is respective type
type PostCache struct {
	TemplPostMap map[string]*model.Post
	TemplStore   []*model.Post
}

func InitCache() (*PostCache, error) {
	posts, err := model.NewPostArray()
	if err != nil {
		return nil, err
	}

	postList := []*model.Post{}

	templPostMap := map[string]*model.Post{}

	for _, p := range posts {

		// could be handled in the template, but then it would allow for its url /slug to be pulled
		// that could be resolved with a handler check and error if slug is not published
		// this seems currently to be a cleaner way to just keep unpublished posts out of the cache
		if p.Published {
			postList = append(postList, p)
		}

		templPostMap[p.Slug] = p
	}

	return &PostCache{
		TemplStore:   postList,
		TemplPostMap: templPostMap,
	}, nil
}

package raindrop

import (
	"github.com/bananazon/raindrop/pkg/api"
	"github.com/bananazon/raindrop/pkg/data"
)

func (r *Raindrop) ListHighlights() (highlights map[string]*data.Highlight, err error) {
	highlights, err = r.getAllHighlights()
	if err != nil {
		return highlights, err
	}

	return highlights, nil
}

func (r *Raindrop) getAllHighlights() (highlights map[string]*data.Highlight, err error) {
	highlights = make(map[string]*data.Highlight)
	page := 0

	for {
		listBookmarksResult, err := r.API.ListHighlights(page)
		if err != nil {
			return highlights, err
		}

		for _, highlight := range listBookmarksResult.Items {
			highlights[highlight.Id] = highlight
		}

		over := len(listBookmarksResult.Items) < api.PageSize

		if over {
			break
		}
		page += 1
	}

	return highlights, nil
}

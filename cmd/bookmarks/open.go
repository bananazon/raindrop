package bookmarks

import (
	"fmt"
	"os"

	"github.com/bananazon/raindrop/pkg/context"
	"github.com/bananazon/raindrop/pkg/data"
	"github.com/bananazon/raindrop/pkg/raindrop"
	"github.com/nexidian/gocliselect"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func paginate(items []*data.Bookmark, pageSize int) [][]*data.Bookmark {
	var pages [][]*data.Bookmark
	for i := 0; i < len(items); i += pageSize {
		end := i + pageSize
		if end > len(items) {
			end = len(items)
		}
		pages = append(pages, items[i:end])
	}
	return pages
}

func newOpenBookmarkCmd(ctx *context.AppContext) (c *cobra.Command) {
	c = &cobra.Command{
		Use:     "open",
		Aliases: []string{"o"},
		Short:   "Open a bookmark from your raindrop.io account",
		PreRun: func(cmdC *cobra.Command, args []string) {
			rd, err := raindrop.New(ctx.RaindropHome, ctx.RaindropConfig, ctx.Logger)
			if err != nil {
				ctx.Logger.Errorf("Failed to initialize raindrop: %s", err.Error())
				ctx.Logger.Exit(1)
			}
			ctx.RD = rd
		},
		Run: func(cmdC *cobra.Command, args []string) {
			bookmarksSlice := []*data.Bookmark{}
			bookmarks, err := ctx.RD.ListBookmarks()
			if err != nil {
				ctx.Logger.Errorf("Failed to get a list of bookmarks: %s", err.Error())
				ctx.Logger.Exit(1)
			}

			collections, err := ctx.RD.ListCollections()
			if err != nil {
				ctx.Logger.Errorf("Failed to get a list of collections: %s", err.Error())
				ctx.Logger.Exit(1)
			}

			for _, bookmark := range bookmarks {
				bookmarksSlice = append(bookmarksSlice, bookmark)
			}

			pageSize := ctx.ScreenHeight - 5
			pages := paginate(bookmarksSlice, pageSize)
			page := 0

			// Find and set collection name
			for idx, _ := range bookmarks {
				bookmarks[idx].Collection.Name = "Unsorted"
				collectionId := bookmarks[idx].Collection.Id
				collection, exists := collections[uint64(collectionId)]
				if exists {
					if bookmarks[idx].Collection.Id > 0 {
						bookmarks[idx].Collection.Name = collection.Title
					}
				}
			}

			for {
				menu := gocliselect.NewMenu("Select a Bookmark (esc to cancel)")

				// Add navigation items if needed
				if page > 0 {
					menu.AddItem("Prev Page", "Prev Page")
				}

				// Add current page items
				for _, b := range pages[page] {
					menu.AddItem(b.Link, b.Link) // label and value
				}

				if page < len(pages)-1 {
					menu.AddItem("Next Page", "Next Page")
				}

				choice := menu.Display()

				switch choice {
				case "Next Page":
					if page < len(pages)-1 {
						page++
					}
				case "Prev Page":
					if page > 0 {
						page--
					}
				case "":
					// Esc pressed, exit
					fmt.Println("\nCancelled")
					os.Exit(0)
				default:
					// Find the selected Bookmark
					var selected *data.Bookmark
					for _, b := range pages[page] {
						if b.Link == choice {
							selected = b
							break
						}
					}

					if selected != nil {
						err := browser.OpenURL(selected.Link)
						if err != nil {
							ctx.Logger.Errorf("Failed to open the browser: %s", err.Error())
							ctx.Logger.Exit(1)
						}
					}
					os.Exit(0)
				}
			}
		},
	}

	return c
}

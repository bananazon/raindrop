package highlights

import (
	"github.com/bananazon/raindrop/pkg/context"
	"github.com/bananazon/raindrop/pkg/data"
	"github.com/bananazon/raindrop/pkg/raindrop"
	"github.com/spf13/cobra"
)

func newRemoveHighlightsCmd(ctx *context.AppContext) (c *cobra.Command) {
	c = &cobra.Command{
		Use:     "remove",
		Aliases: []string{"r"},
		Short:   "Remove existing highlights from your raindrop.io account",
		PreRun: func(cmdC *cobra.Command, args []string) {
			rd, err := raindrop.New(ctx.RaindropHome, ctx.RaindropConfig, ctx.Logger)
			if err != nil {
				ctx.Logger.Errorf("Failed to initialize raindrop: %s", err.Error())
				ctx.Logger.Exit(1)
			}
			ctx.RD = rd
		},
		Run: func(cmdC *cobra.Command, args []string) {
			var highlights []data.RemoveHighlightItem
			for _, highLightId := range ctx.FlagRemoveHighlightHighlightId {
				highlights = append(highlights, data.RemoveHighlightItem{
					Id: highLightId,
				})

			}
			_, err := ctx.RD.API.RemoveHighlight(
				ctx.FlagRemoveHighlightBookmarkId,
				data.RemoveHighlightsPayload{
					Highlights: highlights,
				},
			)
			if err != nil {
				ctx.Logger.Errorf("Failed to remove the bookmark: %s", err.Error())
				ctx.Logger.Exit(1)
			}
			ctx.Logger.Infof("Successfully removed the bookmark")
		},
	}

	ctx.GetRemoveHighlightFlags(c)

	return c
}

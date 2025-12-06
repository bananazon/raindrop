package collections

import (
	"github.com/bananazon/raindrop/pkg/context"
	"github.com/bananazon/raindrop/pkg/raindrop"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

func newListCollectionsChildrenCmd(ctx *context.AppContext) (c *cobra.Command) {
	c = &cobra.Command{
		Use:     "list-children",
		Aliases: []string{"lc"},
		Short:   "List the child collections in your raindrop.io account",
		PreRun: func(cmdC *cobra.Command, args []string) {
			rd, err := raindrop.New(ctx.RaindropHome, ctx.RaindropConfig, ctx.Logger)
			if err != nil {
				ctx.Logger.Errorf("Failed to initialize raindrop: %s", err.Error())
				ctx.Logger.Exit(1)
			}
			ctx.RD = rd
		},
		Run: func(cmdC *cobra.Command, args []string) {
			children, err := ctx.RD.ListCollectionsChildren()
			if err != nil {
				ctx.Logger.Errorf("Failed to get a list of collection children: %s", err.Error())
				ctx.Logger.Exit(1)
			}
			pretty.Println(children)

		},
	}

	return c

}

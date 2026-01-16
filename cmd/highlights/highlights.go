package highlights

import (
	"github.com/bananazon/raindrop/pkg/context"
	"github.com/spf13/cobra"
)

func NewHighlightsCmd(ctx *context.AppContext) (cmdC *cobra.Command) {
	cmdC = &cobra.Command{
		Use:     "highlights",
		Aliases: []string{"h"},
		Short:   "Manage highlights in your raindrop.io account",
		Long:    "Manage highlights in your raindrop.io account",
	}

	cmdC.AddCommand(newListHighlightsCmd(ctx))
	// cmdC.AddCommand(newRemoveHighlightsCmd(ctx))

	return cmdC
}

package highlights

import (
	"fmt"
	"os"

	"github.com/bananazon/raindrop/pkg/context"
	"github.com/bananazon/raindrop/pkg/raindrop"
	rdtable "github.com/bananazon/raindrop/pkg/table"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func newListHighlightsCmd(ctx *context.AppContext) (c *cobra.Command) {
	c = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "ls"},
		Short:   "List the highlights in your raindrop.io account",
		PreRun: func(cmdC *cobra.Command, args []string) {
			rd, err := raindrop.New(ctx.RaindropHome, ctx.RaindropConfig, ctx.Logger)
			if err != nil {
				ctx.Logger.Errorf("Failed to initialize raindrop: %s", err.Error())
				ctx.Logger.Exit(1)
			}
			ctx.RD = rd
		},
		Run: func(cmdC *cobra.Command, args []string) {
			var useList = false

			highlights, err := ctx.RD.ListHighlights()
			if err != nil {
				ctx.Logger.Errorf("Failed to get a list of highlights: %s", err.Error())
				ctx.Logger.Exit(1)
			}

			if ctx.ScreenWidth < 80 {
				ctx.Logger.Infof("Screen width is only %d, using list output mode", ctx.ScreenWidth)
				useList = true
			}

			if ctx.FlagPageStyle == "list" {
				useList = true
			}

			if len(highlights) <= 0 {
				fmt.Fprintln(os.Stdout, "No highlights found")
				return
			}

			if useList {
				for _, highlight := range highlights {
					fmt.Fprintf(os.Stdout, "%s = %s\n", "      id", highlight.Id)
					fmt.Fprintf(os.Stdout, "%s = %s\n", "   title", highlight.Title)
					fmt.Fprintf(os.Stdout, "%s = %s\n", "    text", highlight.Text)
					fmt.Fprintf(os.Stdout, "%s = %s\n", "    note", highlight.Note)
					fmt.Fprintln(os.Stdout, "")
				}
			} else {
				maxIdWidth := 25
				maxLinkWidth := 50
				maxTextWidth := 50

				t := rdtable.GetTableTemplate("Highlights", ctx.FlagPageSize, ctx.FlagPageStyle)

				t.SetColumnConfigs([]table.ColumnConfig{
					{Name: "ID", WidthMax: maxIdWidth},
					{Name: "Link", WidthMax: maxLinkWidth},
					{Name: "Text", WidthMax: maxTextWidth},
					{Name: "Note", WidthMax: ctx.ScreenWidth - (maxLinkWidth + maxTextWidth)},
				})

				t.SortBy([]table.SortBy{{Name: "Title", Mode: table.Asc}})
				t.AppendHeader(table.Row{"ID", "Link", "Text", "Note"})

				for _, highlight := range highlights {
					t.AppendRow(table.Row{
						highlight.Id,
						highlight.Link,
						highlight.Text,
						highlight.Note,
					})
				}

				fmt.Fprintln(os.Stdout, t.Render())
			}
		},
	}

	ctx.GetTableFlags(c)

	return c

}

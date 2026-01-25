package add

import (
	"fmt"

	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/Tkdefender88/booky/internal/repo"
	"github.com/Tkdefender88/booky/internal/repo/generated"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.Flags().StringP("description", "d", "", "description of the bookmark")
	Cmd.Flags().StringP("name", "n", "", "name of the bookmark")
	Cmd.Flags().StringArrayP("tags", "t", []string{}, "tags of the bookmark")
}

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "adds a bookmark",
	Long:  "adds a bookmark",
	Example: `
	booky add https://google.com -d "searching the web"
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("bookmark url is required")
		}

		bookmarkURL := args[0]

		var bookmarkTitle string = "bookmark"
		if flags := cmd.Flags(); flags.Changed("name") {
			bookmarkTitle, _ = flags.GetString("name")
		}

		var bookmarkDescription string = ""
		if flags := cmd.Flags(); flags.Changed("description") {
			bookmarkDescription, _ = flags.GetString("description")
		}

		db, err := repo.NewDB()
		if err != nil {
			return fmt.Errorf("could not open bookmark db: %w", err)
		}
		defer db.Close()

		querier := generated.New(db.DB())
		manager := bookmarks.NewManager(querier)

		ctx := cmd.Context()
		manager.SaveBookmark(ctx, bookmarkTitle, bookmarkURL, bookmarkDescription)

		return nil
	},
}

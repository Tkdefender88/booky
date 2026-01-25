package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Tkdefender88/booky/cmd/add"
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/Tkdefender88/booky/internal/repo"
	"github.com/Tkdefender88/booky/internal/repo/generated"
	"github.com/Tkdefender88/booky/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add.Cmd)
}

var rootCmd = &cobra.Command{
	Use:     "booky",
	Short:   "launches the bookmark view",
	Long:    "search for your bookmarks",
	Example: "booky",
	RunE:    launchTui,
}

func launchTui(cmd *cobra.Command, args []string) error {
	ds, err := repo.NewDB()
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	defer ds.Close()

	querier := generated.New(ds.DB())
	manager := bookmarks.NewManager(querier)

	model := tui.NewModel(manager)
	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

func Execute(ctx context.Context) {
	if err := fang.Execute(ctx, rootCmd); err != nil {
		os.Exit(1)
	}
}

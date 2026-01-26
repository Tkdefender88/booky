package cmd

import (
	"context"
	"os"

	"github.com/Tkdefender88/booky/cmd/add"
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
	model := tui.NewModel()
	p := tea.NewProgram(model)

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

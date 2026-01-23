package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Tkdefender88/booky/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "booky",
	Short:   "launches the bookmark view",
	Long:    "search for your bookmarks",
	Example: "booky",
	RunE:    launchTui,
}

func launchTui(cmd *cobra.Command, args []string) error {
	model := tui.NewModel()
	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

func Execute() {
	ctx := context.Background()
	if err := fang.Execute(ctx, rootCmd); err != nil {
		fmt.Fprintf(os.Stderr, "encountered error running program: %v\n", err)
		os.Exit(1)
	}
}

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "booky",
	Short:   "",
	Long:    "",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	ctx := context.Background()
	if err := fang.Execute(ctx, rootCmd); err != nil {
		fmt.Fprintf(os.Stderr, "encountered error running program: %v\n", err)
		os.Exit(1)
	}
}

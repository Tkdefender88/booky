package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Tkdefender88/booky/cmd"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cmd.Execute(context.Background())
	return nil
}

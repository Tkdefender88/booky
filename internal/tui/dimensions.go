package tui

import (
	"charm.land/lipgloss/v2"
	"github.com/Tkdefender88/booky/internal/tui/styles"
)

// Layout dimension constants
const (
	// Global minimum requirements
	MinHeight = 24
	MinWidth  = 80

	// Minimum content dimensions to ensure usability
	MinContentHeight = 10
	MinContentWidth  = 40
)

// calculateListStyleOverhead measures the actual height overhead of the list styling
// This accounts for borders, padding, and any other styling that adds height
// By measuring actual rendered output, it adapts if the style changes
func calculateListStyleOverhead() int {
	// Create a test-rendered component using the list style
	testContent := styles.ListStyle.Render("test")

	// Measure the actual height of a single-line styled component
	// If there's overhead (borders, padding), Height() will be > 1
	overhead := lipgloss.Height(testContent) - 1

	// Ensure we have at least some reasonable minimum
	if overhead < 0 {
		overhead = 0
	}

	return overhead
}

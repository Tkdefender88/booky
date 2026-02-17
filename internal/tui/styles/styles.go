package styles

import "charm.land/lipgloss/v2"

var (
	Subtle = lipgloss.Color("#383838")
	Active = lipgloss.Color("#ffffff")

	ListStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(Subtle).
			Padding(0, 1)
)

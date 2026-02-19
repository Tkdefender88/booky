package styles

import "charm.land/lipgloss/v2"

var (
	Subtle = lipgloss.Color("#383838")
	Active = lipgloss.Color("#ffffff")

	ListStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(Subtle).
			Padding(0, 1)

	// Modal container with thick border for floating form
	ModalStyle = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(Active).
			Padding(1, 2)

	// Dimmed overlay for background content when modal is open
	DimmedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#606060"))
)

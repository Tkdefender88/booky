package tui

import "charm.land/lipgloss/v2"

var (
	//normal = lipgloss.Color("#eeeeee")
	subtle = lipgloss.Color("#383838")
	active = lipgloss.Color("#ffffff")

	listStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(subtle).
			Padding(0, 1)
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	switch m.state {
	case DBConnecting:
		return m.spinner.View()
	case LoadingBookmarks:
		return m.spinner.View()
	case Success:
		subtle := lipgloss.Color("#383838")
		active := lipgloss.Color("#ffffff")

		tagListStyle := listStyle.
			Width(m.tagList.Width()).
			Height(m.tagList.Height())

		bookmarkListStyle := listStyle.
			Width(m.bookmarkList.Width()).
			Height(m.bookmarkList.Height())

		switch m.focus {
		case bookmarksFocus:
			tagListStyle = tagListStyle.BorderForeground(subtle)
			bookmarkListStyle = bookmarkListStyle.BorderForeground(active)
		case tagsFocus:
			tagListStyle = tagListStyle.BorderForeground(active)
			bookmarkListStyle = bookmarkListStyle.BorderForeground(subtle)
		}

		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			tagListStyle.Render(m.tagList.View()),
			bookmarkListStyle.Render(m.bookmarkList.View()),
		)
	default:
		return ""
	}
}

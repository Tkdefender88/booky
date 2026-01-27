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

	activeListStyle = listStyle.BorderForeground(active)
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
		bookmarkList := ""
		tagList := ""
		switch m.focus {
		case bookmarksFocus:
			tagList = listStyle.Render(m.tagList.View())
			bookmarkList = activeListStyle.Render(m.bookmarkList.View())
		case tagsFocus:
			tagList = activeListStyle.Render(m.tagList.View())
			bookmarkList = listStyle.Render(m.bookmarkList.View())
		default:
			bookmarkList = listStyle.Render(m.bookmarkList.View())
			tagList = listStyle.Render(m.tagList.View())
		}

		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			tagList,
			bookmarkList,
		)
	default:
		return ""
	}
}

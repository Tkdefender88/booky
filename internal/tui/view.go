package tui

import "charm.land/lipgloss/v2"

var (
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
	case DBConnecting, LoadingBookmarks:
		return m.spinner.View()
	case AddingBookmark:
		return m.addBookmark.View()
	case TagsList, BookmarksList:
		tagListStyle := listStyle.
			Width(m.tagList.Width()).
			Height(m.tagList.Height())

		bookmarkListStyle := listStyle.
			Width(m.bookmarkList.Width()).
			Height(m.bookmarkList.Height())

		if m.state == BookmarksList {
			tagListStyle = tagListStyle.BorderForeground(subtle)
			bookmarkListStyle = bookmarkListStyle.BorderForeground(active)
		}
		if m.state == TagsList {
			tagListStyle = tagListStyle.BorderForeground(active)
			bookmarkListStyle = bookmarkListStyle.BorderForeground(subtle)
		}

		help := m.help.View(m.keymap)
		lists := lipgloss.JoinHorizontal(
			lipgloss.Top,
			tagListStyle.Render(m.tagList.View()),
			bookmarkListStyle.Render(m.bookmarkList.View()),
		)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			lists,
			help,
		)
	default:
		return ""
	}
}

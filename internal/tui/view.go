package tui

import (
	"charm.land/lipgloss/v2"
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
		lists := lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.tagList.View(),
			m.bookmarkList.View(),
		)

		// fullView := lipgloss.JoinVertical(
		// 	lipgloss.Left,
		// 	lists,
		// 	"help",
		// )
		return lists
	}

	return ""
}

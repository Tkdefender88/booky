package bookmarklist

import "github.com/Tkdefender88/booky/internal/tui/styles"

func (m Model) View() string {
	bookmarkListStyle := styles.ListStyle.
		Width(m.list.Width()).
		Height(m.list.Height()).
		BorderForeground(styles.Subtle)

	if m.active {
		bookmarkListStyle = bookmarkListStyle.BorderForeground(styles.Active)
	}

	return bookmarkListStyle.Render(m.list.View())
}


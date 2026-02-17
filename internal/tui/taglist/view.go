package taglist

import "github.com/Tkdefender88/booky/internal/tui/styles"

func (m Model) View() string {
	listView := styles.ListStyle.
		Width(m.list.Width()).
		Height(m.list.Height()).
		BorderForeground(styles.Subtle)

	if m.active {
		listView = listView.BorderForeground(styles.Active)
	}

	return listView.Render(m.list.View())
}

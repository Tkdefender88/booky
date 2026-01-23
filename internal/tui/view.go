package tui

func (m Model) View() string {
	switch m.state {
	case Loading:
		return m.spinner.View()
	case Success:
		return m.list.View()
	default:
		return ""
	}
}

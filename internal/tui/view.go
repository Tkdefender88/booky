package tui

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
		return m.list.View()
	default:
		return ""
	}
}

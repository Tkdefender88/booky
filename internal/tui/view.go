package tui

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/Tkdefender88/booky/internal/tui/keys"
	"github.com/charmbracelet/bubbles/key"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	switch m.state {
	case DBConnecting, LoadingBookmarks:
		return m.spinner.View()
	case AddingBookmark:
		// huh form has built-in help display at the bottom
		return m.addBookmark.View()
	case TagsList, BookmarksList:
		lists := lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.tagList.View(),
			m.bookmarkList.View(),
		)

		// Add help footer
		footer := m.helpFooter()

		fullView := lipgloss.JoinVertical(
			lipgloss.Left,
			lists,
			footer,
		)
		return fullView
	}

	return ""
}

// helpFooter generates a context-aware single-line help footer
func (m Model) helpFooter() string {
	var bindings []key.Binding

	// Always show global keys
	bindings = append(bindings, keys.Global.AddBookmark)
	bindings = append(bindings, keys.Global.Quit)

	// Add navigation keys
	bindings = append(bindings, keys.Navigation.LineUp)
	bindings = append(bindings, keys.Navigation.LineDown)
	bindings = append(bindings, keys.Navigation.SwitchView)

	// Add component-specific keys based on active state
	switch m.state {
	case TagsList:
		if kp, ok := any(m.tagList).(KeyProvider); ok {
			bindings = append(bindings, kp.HelpBindings()...)
		}
	case BookmarksList:
		if kp, ok := any(m.bookmarkList).(KeyProvider); ok {
			bindings = append(bindings, kp.HelpBindings()...)
		}
	}

	return formatHelpLine(bindings)
}

// formatHelpLine formats key bindings into a single line
func formatHelpLine(bindings []key.Binding) string {
	var parts []string
	for _, binding := range bindings {
		h := binding.Help()
		parts = append(parts, h.Key+": "+h.Desc)
	}

	line := strings.Join(parts, "  │  ")

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("8")).
		Padding(0, 1)

	return style.Render(line)
}

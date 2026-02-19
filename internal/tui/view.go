package tui

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/Tkdefender88/booky/internal/tui/keys"
	"github.com/Tkdefender88/booky/internal/tui/styles"
	"github.com/charmbracelet/bubbles/key"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	switch m.state {
	case DBConnecting, LoadingBookmarks:
		return m.spinner.View()
	case TagsList, BookmarksList, AddingBookmark:
		// Always render the base lists view
		listsView := m.renderLists()

		// If form is open, layer it as a modal overlay
		if m.state == AddingBookmark {
			return m.renderWithModalOverlay(listsView)
		}

		return listsView
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

// renderLists renders the bookmark and tag lists with help footer
func (m Model) renderLists() string {
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

// renderFormModal renders the form wrapped in a modal style
func (m Model) renderFormModal() string {
	formContent := m.addBookmark.View()
	modal := styles.ModalStyle.Render(formContent)
	return modal
}

// renderWithModalOverlay layers the form modal over the dimmed lists view
func (m Model) renderWithModalOverlay(listsView string) string {
	// Dim the background lists
	dimmedLists := styles.DimmedStyle.Render(listsView)

	// Create the modal
	modal := m.renderFormModal()

	// Get modal dimensions
	modalLines := strings.Split(modal, "\n")
	modalHeight := len(modalLines)
	modalWidth := 0
	for _, line := range modalLines {
		// Use lipgloss Width to account for ANSI codes
		w := lipgloss.Width(line)
		if w > modalWidth {
			modalWidth = w
		}
	}

	// Calculate center position
	x := (m.width - modalWidth) / 2
	y := (m.height - modalHeight) / 2
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	// Create layers using lipgloss Canvas
	backgroundLayer := lipgloss.NewLayer(dimmedLists).X(0).Y(0).Z(0)
	modalLayer := lipgloss.NewLayer(modal).X(x).Y(y).Z(1)

	canvas := lipgloss.NewCanvas(backgroundLayer, modalLayer)

	// Render the canvas to a string
	return canvas.Render()
}

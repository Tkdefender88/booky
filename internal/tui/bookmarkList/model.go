package bookmarklist

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	active bool
	list   list.Model
	width  int
	height int
}

func NewModel() Model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.SetShowHelp(false)

	// Disable the default 'q' quit key binding - we handle quit at the app level
	list.KeyMap.Quit.Unbind()

	return Model{
		list: list,
	}
}

func (m *Model) SetActive(active bool) {
	m.active = active
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.list.SetSize(width, height)
}

func (m Model) Init() tea.Cmd {
	return nil
}

// HelpBindings returns the key bindings for the bookmark list component
// Implements the KeyProvider interface
func (m Model) HelpBindings() []key.Binding {
	return []key.Binding{localKeys.Open}
}

// FilterState returns the current filtering state of the list
// This allows parent components to check if user is actively filtering
func (m Model) FilterState() list.FilterState {
	return m.list.FilterState()
}

// SetFilterState sets the filtering state of the list
// This is primarily used for testing purposes
func (m *Model) SetFilterState(state list.FilterState) {
	m.list.SetFilterState(state)
}

// FilteringEnabled returns whether filtering is enabled
func (m Model) FilteringEnabled() bool {
	return m.list.FilteringEnabled()
}

// SetItems sets the items in the list
func (m *Model) SetItems(items []list.Item) {
	m.list.SetItems(items)
}

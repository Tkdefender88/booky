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

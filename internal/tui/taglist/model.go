package taglist

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list        list.Model
	active      bool
	width       int
	height      int
	manager     *bookmarks.BookmarkManager
	selectedTag string
}

func NewModel() Model {
	list := list.New([]list.Item{}, tagDelegate{}, 0, 0)
	list.Title = "Tags"
	list.SetShowTitle(true)
	list.SetShowHelp(false)

	// Disable the default 'q' quit key binding - we handle quit at the app level
	list.KeyMap.Quit.Unbind()

	return Model{
		active: true,
		list:   list,
	}
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.list.SetSize(width, height)
}

func (m Model) Init() tea.Cmd {
	return nil
}

// HelpBindings returns the key bindings for the tag list component
// Implements the KeyProvider interface
func (m Model) HelpBindings() []key.Binding {
	// Tag list currently has no unique keys
	return []key.Binding{}
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

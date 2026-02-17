package taglist

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list        list.Model
	active      bool
	keymap      KeyMap
	width       int
	height      int
	manager     *bookmarks.BookmarkManager
	selectedTag string
}

func NewModel() Model {
	list := list.New([]list.Item{}, tagDelegate{}, 0, 0)
	list.SetShowHelp(false)
	return Model{
		active: true,
		list:   list,
		keymap: KeyBinds(),
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

package tui

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type State int

type ErrMsg struct {
	err error
}

const (
	DBConnecting State = iota
	LoadingBookmarks
	Success
)

type Model struct {
	list    list.Model
	spinner spinner.Model
	manager *bookmarks.BookmarkManager

	state    State
	shutdown func() error

	err error
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, ConnectDB())
}

func NewModel() Model {
	spinner := spinner.New()
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	return Model{
		spinner: spinner,
		state:   DBConnecting,
		list:    list,
	}
}

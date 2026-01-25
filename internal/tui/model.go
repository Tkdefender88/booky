package tui

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	Loading State = iota
	Success
)

type Model struct {
	list    list.Model
	spinner spinner.Model
	manager *bookmarks.BookmarkManager

	state State
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func NewModel(manager *bookmarks.BookmarkManager) Model {
	items := []list.Item{
		&item{title: "google", url: "https://google.com", desc: "searching the web"},
	}

	bookmarkList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	spinner := spinner.New()
	return Model{
		list:    bookmarkList,
		spinner: spinner,
		manager: manager,

		state: Loading,
	}
}

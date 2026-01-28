package tui

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type State int
type Focus int

func nextFocus(f Focus) Focus {
	return Focus((f + 1) % 2)
}

type ErrMsg struct {
	err error
}

const (
	tagsFocus Focus = iota
	bookmarksFocus
)

const (
	DBConnecting State = iota
	LoadingBookmarks
	Success
)

type Model struct {
	bookmarkList list.Model
	tagList      list.Model
	spinner      spinner.Model
	help         help.Model

	keymap KeyMap

	manager *bookmarks.BookmarkManager

	state       State
	focus       Focus
	shutdown    func() error
	selectedTag string

	err error
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, ConnectDB())
}

func NewModel() Model {
	spinner := spinner.New()
	bmList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	tagList := list.New([]list.Item{}, tagDelegate{}, 0, 0)
	tagList.SetShowHelp(false)
	bmList.SetShowHelp(false)

	return Model{
		spinner:      spinner,
		bookmarkList: bmList,
		tagList:      tagList,
		help:         help.New(),

		keymap: TagsKeyMap(),

		state:       DBConnecting,
		focus:       tagsFocus,
		selectedTag: "",
	}
}

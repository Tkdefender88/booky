package tui

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type State int

type ErrMsg struct {
	err error
}

type FormKey string

const (
	Name        FormKey = "Name"
	Description FormKey = "Description"
	Url         FormKey = "Url"
	Tags        FormKey = "Tags"
)

const (
	DBConnecting State = iota
	LoadingBookmarks
	BookmarksList
	TagsList
	AddingBookmark
)

type Model struct {
	bookmarkList list.Model
	tagList      list.Model
	spinner      spinner.Model
	help         help.Model
	addBookmark  *huh.Form

	keymap KeyMap

	manager *bookmarks.BookmarkManager

	state       State
	shutdown    func() error
	selectedTag string

	err error
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		ConnectDB(),
	)
}

func createForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key(string(Name)).
				Title("Bookmark Name").
				Prompt("? "),
			huh.NewInput().
				Key(string(Url)).
				Title("Url").
				Prompt("? "),
			huh.NewInput().
				Key(string(Description)).
				Title("Description").
				Prompt("? "),
			huh.NewInput().
				Key(string(Tags)).
				Title("Tags").
				Prompt("? "),
		),
	)
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
		addBookmark:  nil,

		keymap: TagsKeyMap(),

		state:       DBConnecting,
		selectedTag: "",
	}
}

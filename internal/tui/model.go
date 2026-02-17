package tui

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	bookmarklist "github.com/Tkdefender88/booky/internal/tui/bookmarkList"
	"github.com/Tkdefender88/booky/internal/tui/taglist"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type State int

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
	state State

	bookmarkList bookmarklist.Model
	tagList      taglist.Model
	spinner      spinner.Model
	help         help.Model
	addBookmark  *huh.Form

	keymap KeyMap

	manager  *bookmarks.BookmarkManager
	shutdown func() error

	err error
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		ConnectDB(),
		m.addBookmark.Init(),
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
	bmList := bookmarklist.NewModel()
	tagList := taglist.NewModel()

	return Model{
		spinner:      spinner,
		bookmarkList: bmList,
		tagList:      tagList,
		help:         help.New(),
		addBookmark:  createForm(),

		keymap: Keymap(),

		state: DBConnecting,
	}
}

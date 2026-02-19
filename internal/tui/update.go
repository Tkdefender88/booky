package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tkdefender88/booky/internal/tui/keys"
	"github.com/Tkdefender88/booky/internal/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/davecgh/go-spew/spew"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.dump != nil {
		spew.Fdump(m.dump, msg)
	}
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case messages.ErrMsg:
		m.err = msg.Err
	case messages.DbConnectedMsg:
		m.state = LoadingBookmarks
		m.manager = msg.Manager
		m.shutdown = msg.Close
		cmds = append(cmds, FetchBookmarks(msg.Manager))
	case messages.BookmarksFetchedMsg:
		m.spinning = false
		m.state = TagsList
	case messages.BookmarkAddedMsg:
		cmds = append(cmds, FetchBookmarks(m.manager))
	case messages.ChangeListFocusMsg:
		switch msg.Target {
		case messages.BookmarkFocus:
			m.state = BookmarksList
		case messages.TagFocus:
			m.state = TagsList
		case messages.FormFocus:
			m.state = AddingBookmark
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		if m.spinning {
			cmds = append(cmds, cmd)
		}
	case messages.FormClosedMsg:
		if msg.Status == messages.FormClosedSuccess {
			cmds = append(cmds, AddBookmark(m.manager, msg.Name, msg.Url, msg.Desc, msg.Tags))
		}
		m.addBookmark = createForm()
		cmds = append(cmds, changeToTagsFocus, m.addBookmark.Init())
	case tea.WindowSizeMsg:
		m = updateWindowSize(m, msg)
	case tea.QuitMsg:
		if m.shutdown != nil {
			if err := m.shutdown(); err != nil {
				fmt.Fprintf(os.Stderr, "error shutting down: %v\n", err)
			}
		}
		os.Exit(0)
	case tea.KeyMsg:
		// Check if user is actively filtering in either list
		isFiltering := m.bookmarkList.FilterState() == list.Filtering ||
			m.tagList.FilterState() == list.Filtering

		// Skip global key bindings while filtering (Ctrl+C always works)
		if !isFiltering {
			switch {
			case key.Matches(msg, keys.Global.AddBookmark):
				if m.state != AddingBookmark {
					cmds = append(cmds, changeToFormFocus())
				}
			case key.Matches(msg, keys.Global.Quit):
				if m.state != AddingBookmark {
					cmds = append(cmds, tea.Quit)
				}
			}
		}
	}

	if m.state == AddingBookmark {
		var formCmd tea.Cmd
		m, formCmd = updateForm(m, msg)
		cmds = append(cmds, formCmd)
	}

	bookmarks, cmd := m.bookmarkList.Update(msg)
	m.bookmarkList = bookmarks
	cmds = append(cmds, cmd)

	tags, cmd := m.tagList.Update(msg)
	m.tagList = tags
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func updateForm(model Model, msg tea.Msg) (Model, tea.Cmd) {
	var form *huh.Form = nil

	fmodel, cmd := model.addBookmark.Update(msg)
	if addForm, ok := fmodel.(*huh.Form); ok {
		form = addForm
	}

	if form.State == huh.StateCompleted {
		name := form.GetString(string(Name))
		url := form.GetString(string(Url))
		desc := form.GetString(string(Description))
		tags := parseTags(form.GetString(string(Tags)))
		return model, formClosedSuccess(name, url, desc, tags)
	}

	if form.State == huh.StateAborted {
		return model, formClosedAborted
	}

	return model, cmd
}

func parseTags(tags string) []string {
	t := strings.TrimSpace(tags)
	rawTagList := strings.Split(t, ",")
	var tagList []string
	for _, tag := range rawTagList {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			tagList = append(tagList, tag)
		}
	}
	return tagList
}

func updateWindowSize(model Model, msg tea.WindowSizeMsg) Model {
	// Store terminal dimensions in model
	model.width = msg.Width
	model.height = msg.Height
	model.help.Width = msg.Width

	// Calculate layout dimensions
	listHeight := model.getListHeight()
	tagWidth, bookmarkWidth := model.getListWidths()

	// Apply dimensions to sub-components
	model.bookmarkList.SetSize(bookmarkWidth, listHeight)
	model.tagList.SetSize(tagWidth, listHeight)

	return model
}

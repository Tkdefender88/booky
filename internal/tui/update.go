package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case AddBookmarkMsg:
		m.state = TagsList
		return m, FetchBookmarks(m.manager)
	case DbConnectMsg:
		m.state = LoadingBookmarks
		m.manager = msg.manager
		m.shutdown = msg.close
		return m, FetchBookmarks(msg.manager)
	case BookmarksMsg:
		return updateLists(m, msg)
	case tea.KeyMsg:
		return handleKey(m, msg)
	case tea.WindowSizeMsg:
		m = updateWindowSize(m, msg)
		return m, nil
	case tea.QuitMsg:
		if m.shutdown != nil {
			if err := m.shutdown(); err != nil {
				fmt.Fprintf(os.Stderr, "error shutting down: %v\n", err)
			}
		}
		os.Exit(0)
	}

	var spinCmd tea.Cmd
	m.spinner, spinCmd = m.spinner.Update(msg)
	cmds = append(cmds, spinCmd)

	var interactionCmd tea.Cmd
	switch m.state {
	case AddingBookmark:
		m, interactionCmd = updateForm(m, msg)
	case BookmarksList:
		m.bookmarkList, interactionCmd = m.bookmarkList.Update(msg)
	case TagsList:
		m.tagList, interactionCmd = m.tagList.Update(msg)
		if selectedTag, ok := m.tagList.SelectedItem().(tag); ok {
			if m.selectedTag != selectedTag.Name() {
				m.selectedTag = selectedTag.Name()
				cmds = append(cmds, FetchBookmarksByTag(selectedTag.name, m.manager))
			}
		}
	default:
	}
	cmds = append(cmds, interactionCmd)

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
		model.state = TagsList
		model.addBookmark = form
		return model, AddBookmark(model.manager, name, url, desc, tags)
	}

	if form.State == huh.StateAborted {
		model.addBookmark = form
		model.state = TagsList
		return model, nil
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

func updateLists(model Model, msg BookmarksMsg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	model.state = TagsList
	bookmarks := make([]list.Item, 0, len(msg.bookmarks))
	for _, b := range msg.bookmarks {
		bookmarks = append(bookmarks, fromBookmark(b))
	}
	cmds = append(cmds, model.bookmarkList.SetItems(bookmarks))

	if len(msg.tags) > 0 {
		tags := make([]list.Item, 0, len(msg.tags))
		for _, t := range msg.tags {
			tags = append(tags, newTag(t))
		}
		cmds = append(cmds, model.tagList.SetItems(tags))
	}

	return model, tea.Batch(cmds...)
}

func handleKey(model Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	if model.state == AddingBookmark {
		return model, nil
	}
	var cmds []tea.Cmd
	switch {
	case key.Matches(msg, model.keymap.Quit):
		cmds = append(cmds, tea.Quit)
	case key.Matches(msg, model.keymap.AddBookmark):
		form := createForm()
		model.addBookmark = form
		cmds = append(cmds, model.addBookmark.Init())
		model.state = AddingBookmark
	case key.Matches(msg, model.keymap.SwitchView):
		if model.state == TagsList && model.tagList.FilterState() == list.Filtering {
			return model, nil
		}
		// Switch state and keymaps
		switch model.state {
		case TagsList:
			model.state = BookmarksList
			model.keymap = BookmarksKeyMap()
		case BookmarksList:
			model.state = TagsList
			model.keymap = TagsKeyMap()
		}
	case key.Matches(msg, model.keymap.Open):
		if model.bookmarkList.FilterState() == list.Filtering ||
			model.tagList.FilterState() == list.Filtering {
			return model, nil
		}

		if bookmark, ok := model.bookmarkList.SelectedItem().(bookmark); ok {
			cmds = append(cmds, openBookmark(bookmark))
		}
	}
	return model, tea.Batch(cmds...)
}

func updateWindowSize(model Model, msg tea.WindowSizeMsg) Model {
	phi := 1.6180
	tagWidth := int(float64(msg.Width) / (phi + 1))
	tagWidth = max(tagWidth, 30)
	bookmarkWidth := msg.Width - tagWidth
	model.help.Width = msg.Width
	model.bookmarkList.SetSize(bookmarkWidth, msg.Height-4)
	model.tagList.SetSize(tagWidth, msg.Height-4)
	return model
}

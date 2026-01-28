package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case DbConnectMsg:
		m.state = LoadingBookmarks
		m.manager = msg.manager
		m.shutdown = msg.close
		cmds = append(cmds, FetchBookmarks(msg.manager))
	case BookmarksMsg:
		var cmd tea.Cmd
		m, cmd = updateLists(m, msg)
		cmds = append(cmds, cmd)
	case tea.KeyMsg:
		var cmd tea.Cmd
		m, cmd = handleKey(m, msg)
		cmds = append(cmds, cmd)
	case tea.WindowSizeMsg:
		m = updateWindowSize(m, msg)
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

	var listCmd tea.Cmd
	switch m.focus {
	case bookmarksFocus:
		m.bookmarkList, listCmd = m.bookmarkList.Update(msg)
	case tagsFocus:
		m.tagList, listCmd = m.tagList.Update(msg)
	}
	cmds = append(cmds, listCmd)

	return m, tea.Batch(cmds...)
}

func updateLists(model Model, msg BookmarksMsg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	model.state = Success
	bookmarks := []list.Item{}
	for _, b := range msg.bookmarks {
		bookmarks = append(bookmarks, fromBookmark(b))
	}
	tags := []list.Item{}
	for _, t := range msg.tags {
		tags = append(tags, newTag(t))
	}
	cmds = append(cmds, model.tagList.SetItems(tags))
	cmds = append(cmds, model.bookmarkList.SetItems(bookmarks))

	return model, tea.Batch(cmds...)
}

func handleKey(model Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch {
	case key.Matches(msg, model.keymap.Quit):
		cmds = append(cmds, tea.Quit)
	case key.Matches(msg, model.keymap.SwitchView):
		model.focus = nextFocus(model.focus)
		// Switch keymaps based on new focus
		if model.focus == bookmarksFocus {
			model.keymap = BookmarksKeyMap()
		} else {
			model.keymap = TagsKeyMap()
		}
	case key.Matches(msg, model.keymap.Open):
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

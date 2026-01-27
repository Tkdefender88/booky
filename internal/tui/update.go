package tui

import (
	"fmt"
	"os"

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
		return m, FetchBookmarks(msg.manager)
	case BookmarksMsg:
		m.state = Success
		bookmarks := []list.Item{}
		for _, b := range msg.bookmarks {
			bookmarks = append(bookmarks, fromBookmark(b))
		}
		tags := []list.Item{}
		for _, t := range msg.tags {
			tags = append(tags, newTag(t))
		}
		cmds = append(cmds, m.tagList.SetItems(tags))
		cmds = append(cmds, m.bookmarkList.SetItems(bookmarks))
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.focus = nextFocus(m.focus)
		}
	case tea.WindowSizeMsg:
		phi := 1.6180
		tagWidth := int(float64(msg.Width) / (phi + 1))
		tagWidth = max(tagWidth, 20)
		bookmarkWidth := msg.Width - tagWidth

		m.bookmarkList.SetSize(bookmarkWidth, msg.Height-2)
		m.tagList.SetSize(tagWidth, msg.Height-2)
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

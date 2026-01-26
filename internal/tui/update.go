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
		items := []list.Item{}
		for _, b := range msg.bookmarks {
			items = append(items, fromBookmark(b))
		}
		return m, m.list.SetItems(items)
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, int(msg.Height/2))
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
	m.list, listCmd = m.list.Update(msg)
	cmds = append(cmds, listCmd)

	return m, tea.Batch(cmds...)
}

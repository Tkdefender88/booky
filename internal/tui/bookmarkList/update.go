package bookmarklist

import (
	"github.com/Tkdefender88/booky/internal/tui/keys"
	"github.com/Tkdefender88/booky/internal/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case messages.ChangeListFocusMsg:
		m.active = msg.Target == messages.BookmarkFocus
	case messages.BookmarksFetchedMsg:
		bookmarks := make([]list.Item, 0, len(msg.Bookmarks))
		for _, b := range msg.Bookmarks {
			bookmarks = append(bookmarks, fromBookmark(b))
		}
		m.list.SetItems(bookmarks)
	case tea.KeyMsg:
		if m.active {
			switch {
			case key.Matches(msg, localKeys.Open):
				if bookmark, ok := m.list.SelectedItem().(bookmark); ok {
					cmds = append(cmds, openBookmark(bookmark))
				}
			case key.Matches(msg, keys.Navigation.SwitchView):
				cmds = append(cmds, switchToTags)
			default:
				list, cmd := m.list.Update(msg)
				m.list = list
				cmds = append(cmds, cmd)
			}
		}
	default:
		// Always pass other messages (like FilterMatchesMsg) to the list
		list, cmd := m.list.Update(msg)
		m.list = list
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

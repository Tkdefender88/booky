package taglist

import (
	"github.com/Tkdefender88/booky/internal/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case messages.DbConnectedMsg:
		m.manager = msg.Manager
	case messages.ChangeListFocusMsg:
		m.active = msg.Target == messages.TagFocus
	case messages.BookmarksFetchedMsg:
		if len(msg.Tags) != 0 {
			tags := make([]list.Item, 0, len(msg.Tags))
			for _, t := range msg.Tags {
				tags = append(tags, newTag(t))
			}
			m.list.SetItems(tags)
		}
	case tea.KeyMsg:
		if m.active {
			switch {
			case key.Matches(msg, m.keymap.SwitchView):
				cmds = append(cmds, switchToBookmarks)
			default:
				list, cmd := m.list.Update(msg)
				if selected, ok := list.SelectedItem().(tag); ok {
					tagName := selected.Name()
					cmds = append(cmds, FetchBookmarksByTag(tagName, m.manager))
				}
				cmds = append(cmds, cmd)
				m.list = list
			}
		}
	}
	return m, tea.Batch(cmds...)
}


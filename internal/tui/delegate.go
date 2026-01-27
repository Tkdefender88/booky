package tui

import (
	"fmt"
	"io"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var _ list.Item = &bookmark{}

type bookmark struct {
	title, desc, url string
}

func (i bookmark) Url() string         { return i.url }
func (i bookmark) Title() string       { return i.title }
func (i bookmark) Description() string { return i.desc }

// FilterValue implements list.Item.
func (i bookmark) FilterValue() string { return i.title }

func fromBookmark(b bookmarks.Bookmark) bookmark {
	return bookmark{
		title: b.Title,
		desc:  b.Description,
		url:   b.Url.String(),
	}
}

type tag struct {
	name string
}

func (i tag) Name() string        { return i.name }
func (i tag) FilterValue() string { return i.name }

func newTag(name string) tag {
	return tag{name: name}
}

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

type tagDelegate struct{}

func (d tagDelegate) Height() int                             { return 1 }
func (d tagDelegate) Spacing() int                            { return 0 }
func (d tagDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d tagDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(tag)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i.Name())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("| " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

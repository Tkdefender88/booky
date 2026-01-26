package tui

import (
	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/charmbracelet/bubbles/list"
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

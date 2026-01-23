package tui

import (
	"github.com/charmbracelet/bubbles/list"
)

var _ list.Item = &item{}

type item struct {
	title, desc, url string
}

func (i *item) Url() string         { return i.url }
func (i *item) Title() string       { return i.title }
func (i *item) Description() string { return i.desc }

// FilterValue implements list.Item.
func (i *item) FilterValue() string { return i.title }

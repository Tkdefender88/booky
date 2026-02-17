package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	// Navigation
	AddBookmark key.Binding
	Quit        key.Binding
}

func Keymap() KeyMap {
	return KeyMap{
		AddBookmark: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add new bookmark"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("ctrl+c", "quit"),
		),
	}
}

package keys

import "github.com/charmbracelet/bubbles/key"

type global struct {
	AddBookmark key.Binding
	Quit        key.Binding
}

// Global contains application-wide key bindings
var Global = global{
	AddBookmark: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add bookmark"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

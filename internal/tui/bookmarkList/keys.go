package bookmarklist

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Open key.Binding
}

// localKeys contains keys specific to the bookmark list component
var localKeys = keyMap{
	Open: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "open bookmark"),
	),
}

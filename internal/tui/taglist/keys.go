package taglist

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	// Future: Add tag-list-specific keys here if needed
	SwitchView key.Binding
}

// localKeys contains keys specific to the tag list component
var localKeys = keyMap{
	SwitchView: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "switch to bookmarks"),
	),
}

package bookmarklist

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	// Actions
	Open       key.Binding
	SwitchView key.Binding
}

func KeyBinds() KeyMap {
	return KeyMap{
		Open: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "open bookmark"),
		),
		SwitchView: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch view"),
		),
	}
}


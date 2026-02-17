package taglist

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	SwitchView key.Binding
}

func KeyBinds() KeyMap {
	return KeyMap{
		SwitchView: key.NewBinding(
			key.WithKeys("tab", "enter"),
			key.WithHelp("tab", "switch view"),
		),
	}
}

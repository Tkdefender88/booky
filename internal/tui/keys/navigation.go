package keys

import "github.com/charmbracelet/bubbles/key"

type navigation struct {
	LineUp     key.Binding
	LineDown   key.Binding
	SwitchView key.Binding
}

// Navigation contains navigation-related key bindings
var Navigation = navigation{
	LineUp: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	LineDown: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	SwitchView: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch view"),
	),
}

package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	// Navigation
	SwitchView key.Binding
	Up         key.Binding
	Down       key.Binding
	PageUp     key.Binding
	PageDown   key.Binding

	// Filtering
	Filter key.Binding

	// Actions (optional, only used in bookmarks view)
	Open key.Binding
	Quit key.Binding
}

func CommonKeyBindings() map[string]key.Binding {
	return map[string]key.Binding{
		"switchView": key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch view"),
		),
		"up": key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		"down": key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		"pageUp": key.NewBinding(
			key.WithKeys("pgup"),
			key.WithHelp("pgup", "page up"),
		),
		"pageDown": key.NewBinding(
			key.WithKeys("pgdn"),
			key.WithHelp("pgdn", "page down"),
		),
		"filter": key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "filter"),
		),
		"quit": key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("ctrl+c", "quit"),
		),
	}
}

func TagsKeyMap() KeyMap {
	common := CommonKeyBindings()
	return KeyMap{
		SwitchView: key.NewBinding(
			key.WithKeys("tab", "enter"),
			key.WithHelp("tab/enter", "switch view"),
		),
		Up:       common["up"],
		Down:     common["down"],
		PageUp:   common["pageUp"],
		PageDown: common["pageDown"],
		Filter:   common["filter"],
		Quit:     common["quit"],
	}
}

func BookmarksKeyMap() KeyMap {
	common := CommonKeyBindings()
	return KeyMap{
		SwitchView: common["switchView"],
		Up:         common["up"],
		Down:       common["down"],
		PageUp:     common["pageUp"],
		PageDown:   common["pageDown"],
		Filter:     common["filter"],
		Open: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "open bookmark"),
		),
		Quit: common["quit"],
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	bindings := []key.Binding{k.SwitchView, k.Up, k.Down, k.Filter}
	if k.Open.Help().Key != "" {
		bindings = append(bindings, k.Open)
	}
	bindings = append(bindings, k.Quit)
	return bindings
}

func (k KeyMap) FullHelp() [][]key.Binding {
	help := [][]key.Binding{
		{k.SwitchView, k.Up, k.Down},
		{k.PageUp, k.PageDown, k.Filter},
	}
	if k.Open.Help().Key != "" {
		help = append(help, []key.Binding{k.Open})
	}
	help = append(help, []key.Binding{k.Quit})
	return help
}

package tui

import (
	"testing"

	"github.com/Tkdefender88/booky/internal/tui/keys"
	"github.com/Tkdefender88/booky/internal/tui/messages"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// TestGlobalKeyBindingsBlockedDuringFiltering tests that global key bindings
// like 'a' (add bookmark) are blocked when user is actively filtering
func TestGlobalKeyBindingsBlockedDuringFiltering(t *testing.T) {
	tests := []struct {
		name          string
		filterState   list.FilterState
		keyToPress    string
		shouldTrigger bool
		description   string
	}{
		{
			name:          "Add bookmark works when not filtering",
			filterState:   list.Unfiltered,
			keyToPress:    "a",
			shouldTrigger: true,
			description:   "Pressing 'a' should open form when not filtering",
		},
		{
			name:          "Add bookmark blocked while filtering",
			filterState:   list.Filtering,
			keyToPress:    "a",
			shouldTrigger: false,
			description:   "Pressing 'a' should NOT open form while filtering",
		},
		{
			name:          "Add bookmark works with filter applied",
			filterState:   list.FilterApplied,
			keyToPress:    "a",
			shouldTrigger: true,
			description:   "Pressing 'a' should work when filter is applied but not being edited",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a minimal model
			m, err := NewModel(false)
			if err != nil {
				t.Fatalf("Failed to create model: %v", err)
			}

			// Set the filter state on the tag list (which is active by default)
			m.tagList.SetFilterState(tt.filterState)

			// Track initial state
			initialState := m.state

			// Simulate pressing the key
			keyMsg := tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune(tt.keyToPress),
			}

			// Update the model
			_, cmd := m.Update(keyMsg)

			// Execute the command to get the resulting message
			var stateChanged bool
			if cmd != nil {
				msg := cmd()
				// Check if we got a ChangeListFocusMsg for the form
				if focusMsg, ok := msg.(messages.ChangeListFocusMsg); ok && focusMsg.Target == messages.FormFocus {
					stateChanged = true
				}
			}

			// Verify expectations
			if tt.shouldTrigger && !stateChanged {
				t.Errorf("%s: Expected state change but none occurred", tt.description)
			}
			if !tt.shouldTrigger && stateChanged {
				t.Errorf("%s: State changed when it shouldn't have", tt.description)
			}

			// Verify model state didn't change inappropriately
			if !tt.shouldTrigger && m.state != initialState {
				t.Errorf("Model state changed from %v to %v when it shouldn't", initialState, m.state)
			}
		})
	}
}

// TestFilteringInBookmarkList tests filtering behavior in the bookmark list
func TestFilteringInBookmarkList(t *testing.T) {
	m, err := NewModel(false)
	if err != nil {
		t.Fatalf("Failed to create model: %v", err)
	}

	// Set bookmark list as active
	m.bookmarkList.SetActive(true)
	m.state = BookmarksList

	// Set to filtering state
	m.bookmarkList.SetFilterState(list.Filtering)

	// Try to press 'a' while filtering
	keyMsg := tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'a'},
	}

	initialState := m.state
	_, cmd := m.Update(keyMsg)

	// Should not trigger state change
	if cmd != nil {
		msg := cmd()
		if focusMsg, ok := msg.(messages.ChangeListFocusMsg); ok && focusMsg.Target == messages.FormFocus {
			t.Error("AddBookmark was triggered while filtering in bookmark list")
		}
	}

	if m.state != initialState {
		t.Errorf("State changed from %v to %v while filtering", initialState, m.state)
	}
}

// TestFilteringInTagList tests filtering behavior in the tag list
func TestFilteringInTagList(t *testing.T) {
	m, err := NewModel(false)
	if err != nil {
		t.Fatalf("Failed to create model: %v", err)
	}

	// Set tag list as active (default state)
	m.state = TagsList

	// Set to filtering state
	m.tagList.SetFilterState(list.Filtering)

	// Try to press 'a' while filtering
	keyMsg := tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'a'},
	}

	initialState := m.state
	_, cmd := m.Update(keyMsg)

	// Should not trigger state change
	if cmd != nil {
		msg := cmd()
		if focusMsg, ok := msg.(messages.ChangeListFocusMsg); ok && focusMsg.Target == messages.FormFocus {
			t.Error("AddBookmark was triggered while filtering in tag list")
		}
	}

	if m.state != initialState {
		t.Errorf("State changed from %v to %v while filtering", initialState, m.state)
	}
}

// TestCtrlCAlwaysWorks tests that Ctrl+C works in all states
func TestCtrlCAlwaysWorks(t *testing.T) {
	tests := []struct {
		name        string
		modelState  State
		filterState list.FilterState
	}{
		{"Not filtering", TagsList, list.Unfiltered},
		{"While filtering", TagsList, list.Filtering},
		{"Filter applied", TagsList, list.FilterApplied},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewModel(false)
			if err != nil {
				t.Fatalf("Failed to create model: %v", err)
			}

			m.state = tt.modelState
			m.tagList.SetFilterState(tt.filterState)

			// Simulate Ctrl+C
			keyMsg := tea.KeyMsg{
				Type: tea.KeyCtrlC,
			}

			_, cmd := m.Update(keyMsg)

			// Should always get quit command (might be in a batch)
			if cmd == nil {
				t.Error("Expected quit command but got nil")
			} else {
				msg := cmd()
				// Check if it's a QuitMsg directly or in a BatchMsg
				foundQuit := false
				if _, ok := msg.(tea.QuitMsg); ok {
					foundQuit = true
				}
				if batchMsg, ok := msg.(tea.BatchMsg); ok {
					for _, batchCmd := range batchMsg {
						if batchCmd != nil {
							if _, ok := batchCmd().(tea.QuitMsg); ok {
								foundQuit = true
								break
							}
						}
					}
				}
				if !foundQuit {
					t.Errorf("Expected QuitMsg but got %T", msg)
				}
			}
		})
	}
}

// TestKeyBindingMatchesCorrectly verifies the key binding system works
func TestKeyBindingMatchesCorrectly(t *testing.T) {
	// Test that 'a' matches AddBookmark
	aKey := tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'a'},
	}
	if !key.Matches(aKey, keys.Global.AddBookmark) {
		t.Error("'a' key should match AddBookmark binding")
	}

	// Test that Ctrl+C matches Quit
	ctrlC := tea.KeyMsg{
		Type: tea.KeyCtrlC,
	}
	if !key.Matches(ctrlC, keys.Global.Quit) {
		t.Error("Ctrl+C should match Quit binding")
	}

	// Test that 'q' does NOT match Quit (after our changes)
	qKey := tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'q'},
	}
	if key.Matches(qKey, keys.Global.Quit) {
		t.Error("'q' should NOT match Quit binding after changes")
	}
}

// TestFilterStateExposure verifies that FilterState() methods work correctly
func TestFilterStateExposure(t *testing.T) {
	m, err := NewModel(false)
	if err != nil {
		t.Fatalf("Failed to create model: %v", err)
	}

	// Test bookmark list FilterState exposure
	m.bookmarkList.SetFilterState(list.Filtering)
	if m.bookmarkList.FilterState() != list.Filtering {
		t.Error("BookmarkList FilterState() should return Filtering")
	}

	// Test tag list FilterState exposure
	m.tagList.SetFilterState(list.FilterApplied)
	if m.tagList.FilterState() != list.FilterApplied {
		t.Error("TagList FilterState() should return FilterApplied")
	}

	// Test unfiltered state
	m.tagList.SetFilterState(list.Unfiltered)
	if m.tagList.FilterState() != list.Unfiltered {
		t.Error("TagList FilterState() should return Unfiltered")
	}
}

// TestQKeyDoesNotQuit verifies that pressing 'q' does not quit the application
func TestQKeyDoesNotQuit(t *testing.T) {
	m, err := NewModel(false)
	if err != nil {
		t.Fatalf("Failed to create model: %v", err)
	}

	// Set to normal state (not filtering, not in form)
	m.state = TagsList

	// Simulate pressing 'q'
	keyMsg := tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'q'},
	}

	initialState := m.state
	_, cmd := m.Update(keyMsg)

	// Should NOT generate a quit command
	if cmd != nil {
		msg := cmd()
		// Check if it's a QuitMsg directly or in a BatchMsg
		foundQuit := false
		if _, ok := msg.(tea.QuitMsg); ok {
			foundQuit = true
		}
		if batchMsg, ok := msg.(tea.BatchMsg); ok {
			for _, batchCmd := range batchMsg {
				if batchCmd != nil {
					if _, ok := batchCmd().(tea.QuitMsg); ok {
						foundQuit = true
						break
					}
				}
			}
		}
		if foundQuit {
			t.Error("Pressing 'q' should NOT quit the application")
		}
	}

	// State should remain unchanged
	if m.state != initialState {
		t.Errorf("State changed from %v to %v when pressing 'q'", initialState, m.state)
	}
}

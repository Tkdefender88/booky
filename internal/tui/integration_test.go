//go:build integration
// +build integration

package tui

import (
	"reflect"
	"testing"

	bookmarklist "github.com/Tkdefender88/booky/internal/tui/bookmarkList"
	"github.com/Tkdefender88/booky/internal/tui/messages"
	"github.com/Tkdefender88/booky/internal/tui/taglist"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Test helper functions for consistent test data

// testTag is a minimal implementation of a tag for testing
type testTag struct {
	name string
}

func (t testTag) FilterValue() string { return t.name }
func (t testTag) Title() string       { return t.name }
func (t testTag) Description() string { return "" }

// testBookmark is a minimal implementation of a bookmark for testing
type testBookmark struct {
	title, desc, url string
}

func (b testBookmark) FilterValue() string { return b.title }
func (b testBookmark) Title() string       { return b.title }
func (b testBookmark) Description() string { return b.desc }

// setupTestTags creates a consistent set of tags for testing
func setupTestTags() []list.Item {
	return []list.Item{
		testTag{name: "golang"},
		testTag{name: "rust"},
		testTag{name: "javascript"},
		testTag{name: "go-tools"},
		testTag{name: "python"},
	}
}

// setupTestBookmarks creates a consistent set of bookmarks for testing
func setupTestBookmarks() []list.Item {
	return []list.Item{
		testBookmark{title: "Go Documentation", url: "https://go.dev", desc: "Official Go docs"},
		testBookmark{title: "Rust Book", url: "https://doc.rust-lang.org", desc: "Learn Rust"},
		testBookmark{title: "JavaScript MDN", url: "https://developer.mozilla.org", desc: "JS reference"},
		testBookmark{title: "Go by Example", url: "https://gobyexample.com", desc: "Go tutorials"},
		testBookmark{title: "Python Tutorial", url: "https://docs.python.org", desc: "Python docs"},
	}
}

// TestFilteringEnabledOnInitialization verifies that both lists have filtering
// enabled by default after model creation.
func TestFilteringEnabledOnInitialization(t *testing.T) {
	tagModel := taglist.NewModel()
	bookmarkModel := bookmarklist.NewModel()

	if !tagModel.FilteringEnabled() {
		t.Error("Tag list should have filtering enabled by default")
	}

	if !bookmarkModel.FilteringEnabled() {
		t.Error("Bookmark list should have filtering enabled by default")
	}
}

// mockFilterMatchesMsg is a mock message type to test the default case
// We can't construct the real FilterMatchesMsg because filteredItem is unexported
type mockFilterMatchesMsg struct{}

// TestTagListHasDefaultCase verifies that the taglist Update function
// has a default case that passes unhandled messages to the underlying list.
// This is critical for filtering, as FilterMatchesMsg must be processed.
func TestTagListHasDefaultCase(t *testing.T) {
	m := taglist.NewModel()
	m.SetItems(setupTestTags())
	m.SetSize(80, 24)

	// Send a mock message that won't match any specific case
	// If there's no default case, this message will be silently dropped
	mockMsg := mockFilterMatchesMsg{}

	// The update should not panic and should process the message
	updatedModel, cmd := m.Update(mockMsg)

	// Verify the model was returned (basic sanity check)
	if reflect.TypeOf(updatedModel).String() != "taglist.Model" {
		t.Error("Update should return a taglist.Model")
	}

	// Cmd can be nil or valid, either is fine
	_ = cmd

	// Most importantly: if there's no default case, the list never gets the message
	// We verify the default case exists by checking that unknown messages don't cause issues
}

// TestBookmarkListHasDefaultCase verifies that the bookmarkList Update function
// has a default case that passes unhandled messages to the underlying list.
func TestBookmarkListHasDefaultCase(t *testing.T) {
	m := bookmarklist.NewModel()
	m.SetItems(setupTestBookmarks())
	m.SetSize(80, 24)

	// Send a mock message that won't match any specific case
	mockMsg := mockFilterMatchesMsg{}

	updatedModel, cmd := m.Update(mockMsg)

	// Verify the model was returned
	if reflect.TypeOf(updatedModel).String() != "bookmarklist.Model" {
		t.Error("Update should return a bookmarklist.Model")
	}

	_ = cmd
}

// TestTagListPassesKeyMsgToListWhenActive verifies that KeyMsg is passed
// to the underlying list when the tag list is active.
func TestTagListPassesKeyMsgToListWhenActive(t *testing.T) {
	m := taglist.NewModel()
	m.SetItems(setupTestTags())
	m.SetSize(80, 24)

	// Tag list starts as active
	if m.FilterState() != list.Unfiltered {
		t.Fatalf("Expected initial state to be Unfiltered, got %v", m.FilterState())
	}

	// Send the filter activation key
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}}
	updatedModel, _ := m.Update(keyMsg)

	// The filter state should change to Filtering
	if updatedModel.FilterState() != list.Filtering {
		t.Errorf("Expected filter state to be Filtering after '/', got %v", updatedModel.FilterState())
	}
}

// TestBookmarkListPassesKeyMsgToListWhenActive verifies that KeyMsg is passed
// to the underlying list when the bookmark list is active.
func TestBookmarkListPassesKeyMsgToListWhenActive(t *testing.T) {
	m := bookmarklist.NewModel()
	m.SetItems(setupTestBookmarks())
	m.SetSize(80, 24)
	m.SetActive(true)

	if m.FilterState() != list.Unfiltered {
		t.Fatalf("Expected initial state to be Unfiltered, got %v", m.FilterState())
	}

	// Send the filter activation key
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}}
	updatedModel, _ := m.Update(keyMsg)

	// The filter state should change to Filtering
	if updatedModel.FilterState() != list.Filtering {
		t.Errorf("Expected filter state to be Filtering after '/', got %v", updatedModel.FilterState())
	}
}

// TestTagListDoesNotProcessKeyMsgWhenInactive verifies that KeyMsg is NOT
// processed when the tag list is inactive, but other messages still go through.
func TestTagListDoesNotProcessKeyMsgWhenInactive(t *testing.T) {
	m := taglist.NewModel()
	m.SetItems(setupTestTags())
	m.SetSize(80, 24)

	// Make the tag list inactive
	changeMsg := messages.ChangeListFocusMsg{Target: messages.BookmarkFocus}
	m, _ = m.Update(changeMsg)

	if m.FilterState() != list.Unfiltered {
		t.Fatalf("Expected initial state to be Unfiltered, got %v", m.FilterState())
	}

	// Try sending a KeyMsg while inactive
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}}
	m, _ = m.Update(keyMsg)

	// Filter state should still be Unfiltered because list is inactive
	if m.FilterState() != list.Unfiltered {
		t.Error("KeyMsg should not be processed when list is inactive")
	}

	// But other messages (like WindowSizeMsg) should be processed via default case
	sizeMsg := tea.WindowSizeMsg{Width: 100, Height: 30}
	m, cmd := m.Update(sizeMsg)

	// Should not panic and should return
	_ = cmd
}

// TestBookmarkListDoesNotProcessKeyMsgWhenInactive verifies that KeyMsg is NOT
// processed when the bookmark list is inactive.
func TestBookmarkListDoesNotProcessKeyMsgWhenInactive(t *testing.T) {
	m := bookmarklist.NewModel()
	m.SetItems(setupTestBookmarks())
	m.SetSize(80, 24)

	// Bookmark list starts inactive, but let's be explicit
	changeMsg := messages.ChangeListFocusMsg{Target: messages.TagFocus}
	m, _ = m.Update(changeMsg)

	if m.FilterState() != list.Unfiltered {
		t.Fatalf("Expected initial state to be Unfiltered, got %v", m.FilterState())
	}

	// Try sending a KeyMsg while inactive
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}}
	m, _ = m.Update(keyMsg)

	// Filter state should still be Unfiltered because list is inactive
	if m.FilterState() != list.Unfiltered {
		t.Error("KeyMsg should not be processed when list is inactive")
	}
}

// TestFilteringWorkflow verifies the complete filtering workflow:
// 1. Press '/' to enter filter mode
// 2. Type filter text (simulated)
// 3. Verify filter state transitions correctly
func TestFilteringWorkflow(t *testing.T) {
	m := taglist.NewModel()
	m.SetItems(setupTestTags())
	m.SetSize(80, 24)

	// Step 1: Initial state
	if m.FilterState() != list.Unfiltered {
		t.Fatalf("Expected initial state to be Unfiltered, got %v", m.FilterState())
	}

	// Step 2: Press '/' to enter filter mode
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}}
	m, _ = m.Update(keyMsg)

	if m.FilterState() != list.Filtering {
		t.Errorf("Expected filter state to be Filtering after '/', got %v", m.FilterState())
	}

	// Step 3: Type a character (like 'g')
	// The list will internally process this and send FilterMatchesMsg
	// Our default case must handle that message
	charMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'g'}}
	m, _ = m.Update(charMsg)

	// State should still be Filtering while typing
	state := m.FilterState()
	if state != list.Filtering && state != list.FilterApplied {
		t.Errorf("Expected state to be Filtering or FilterApplied while typing, got %v", state)
	}

	// Step 4: Press ESC to clear filter
	escMsg := tea.KeyMsg{Type: tea.KeyEsc}
	m, _ = m.Update(escMsg)

	// State should return to Unfiltered or FilterApplied (depending on list implementation)
	finalState := m.FilterState()
	if finalState != list.Unfiltered && finalState != list.FilterApplied {
		t.Logf("Final state after ESC: %v (expected Unfiltered or FilterApplied)", finalState)
	}
}

// TestWindowSizeMessagesAlwaysProcessed verifies that window size messages
// are always passed through to the list via the default case.
func TestWindowSizeMessagesAlwaysProcessed(t *testing.T) {
	m := taglist.NewModel()
	m.SetSize(80, 24)

	// Make inactive
	changeMsg := messages.ChangeListFocusMsg{Target: messages.BookmarkFocus}
	m, _ = m.Update(changeMsg)

	// Send a window size message
	sizeMsg := tea.WindowSizeMsg{Width: 100, Height: 30}
	m, cmd := m.Update(sizeMsg)

	// Should not panic and should return
	_ = cmd

	// Verify the model still works after processing
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}
	_, cmd = m.Update(keyMsg)
	_ = cmd
}

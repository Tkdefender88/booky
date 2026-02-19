package taglist

import (
	"context"
	"fmt"

	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/Tkdefender88/booky/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func switchToBookmarks() tea.Msg {
	return messages.ChangeListFocusMsg{Target: messages.BookmarkFocus}
}

func FetchBookmarksByTag(tag string, manager *bookmarks.BookmarkManager) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()

		bookmarks, err := manager.ListBookmarksByTag(ctx, tag)
		if err != nil {
			return messages.NewErrMsg(fmt.Errorf("failed to fetch bookmarks: %w", err))
		}

		return messages.BookmarksFetchedMsg{Bookmarks: bookmarks, Tags: []string{}}
	}
}


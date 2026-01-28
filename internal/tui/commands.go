package tui

import (
	"context"
	"fmt"

	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/Tkdefender88/booky/internal/repo"
	"github.com/Tkdefender88/booky/internal/repo/generated"
	tea "github.com/charmbracelet/bubbletea"
)

type DbConnectMsg struct {
	manager *bookmarks.BookmarkManager
	close   func() error
}

func ConnectDB() tea.Cmd {
	return func() tea.Msg {
		db, err := repo.NewDB()
		if err != nil {
			return ErrMsg{err: fmt.Errorf("failed to open db: %w", err)}
		}

		querier := generated.New(db.DB())
		manager := bookmarks.NewManager(querier)

		return DbConnectMsg{
			manager: manager,
			close:   db.Close,
		}
	}
}

type BookmarksMsg struct {
	bookmarks []bookmarks.Bookmark
	tags      []string
}

func FetchBookmarks(manager *bookmarks.BookmarkManager) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		bookmarks, err := manager.ListBookmarks(ctx)
		if err != nil {
			return ErrMsg{err: fmt.Errorf("failed to fetch bookmarks: %w", err)}
		}

		tags, err := manager.ListTags(ctx)
		if err != nil {
			return ErrMsg{err: fmt.Errorf("failed to fetch tags: %w", err)}
		}
		return BookmarksMsg{bookmarks: bookmarks, tags: tags}
	}
}

func openBookmark(b bookmark) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

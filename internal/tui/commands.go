package tui

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"

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

func FetchBookmarksByTag(tag string, manager *bookmarks.BookmarkManager) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()

		bookmarks, err := manager.ListBookmarksByTag(ctx, tag)
		if err != nil {
			return ErrMsg{err: fmt.Errorf("failed to fetch bookmarks: %w", err)}
		}

		return BookmarksMsg{bookmarks: bookmarks, tags: []string{}}
	}
}

type OpenBrowserMsg struct {
	url string
}

func openBookmark(b bookmark) tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "linux":
			cmd = exec.Command("xdg-open", b.Url())
		case "darwin":
			cmd = exec.Command("open", b.Url())
		default:
			return ErrMsg{err: fmt.Errorf("unsupported platform")}
		}

		if err := cmd.Start(); err != nil {
			return ErrMsg{err: fmt.Errorf("failed to open browser: %w", err)}
		}
		return nil
	}
}

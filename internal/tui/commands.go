package tui

import (
	"context"
	"fmt"

	"github.com/Tkdefender88/booky/internal/bookmarks"
	"github.com/Tkdefender88/booky/internal/repo"
	"github.com/Tkdefender88/booky/internal/repo/generated"
	"github.com/Tkdefender88/booky/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func changeToTagsFocus() tea.Msg {
	return messages.ChangeListFocusMsg{Target: messages.TagFocus}
}

func changeToFormFocus() tea.Cmd {
	return func() tea.Msg {
		return messages.ChangeListFocusMsg{Target: messages.FormFocus}
	}
}

func closeForm(status messages.FormClosedStatus) tea.Cmd {
	return func() tea.Msg {
		return messages.FormClosedMsg{Status: status}
	}
}

func ConnectDB() tea.Cmd {
	return func() tea.Msg {
		db, err := repo.NewDB()
		if err != nil {
			return messages.NewErrMsg(fmt.Errorf("failed to open db: %w", err))
		}

		querier := generated.New(db.DB())
		manager := bookmarks.NewManager(querier)

		return messages.DbConnectedMsg{
			Manager: manager,
			Close:   db.Close,
		}
	}
}

func FetchBookmarks(manager *bookmarks.BookmarkManager) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		bookmarks, err := manager.ListBookmarks(ctx)
		if err != nil {
			return messages.NewErrMsg(fmt.Errorf("failed to fetch bookmarks: %w", err))
		}

		tags, err := manager.ListTags(ctx)
		if err != nil {
			return messages.NewErrMsg(fmt.Errorf("failed to fetch tags: %w", err))
		}
		return messages.BookmarksFetchedMsg{Bookmarks: bookmarks, Tags: tags}
	}
}

func AddBookmark(
	manager *bookmarks.BookmarkManager,
	name, url, desc string,
	tags []string,
) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()

		_, err := manager.SaveBookmark(ctx, name, url, desc, tags)
		if err != nil {
			return messages.NewErrMsg(fmt.Errorf("failed to save bookmark: %w", err))
		}

		return messages.BookmarkAddedMsg{}
	}
}

package bookmarks

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/Tkdefender88/booky/internal/repo/generated"
)

type BookmarkManager struct {
	repo generated.Querier
}

func NewManager(repo generated.Querier) *BookmarkManager {
	return &BookmarkManager{repo: repo}
}

type Bookmark struct {
	Title       string
	Url         *url.URL
	Description string
}

func (m *BookmarkManager) SaveBookmark(ctx context.Context, title, uri, description string) (Bookmark, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return Bookmark{}, fmt.Errorf("failed to parse url: %w", err)
	}

	if err := m.repo.CreateBookmark(ctx, generated.CreateBookmarkParams{
		Title:       title,
		Url:         u.String(),
		Description: sql.NullString{String: description, Valid: true},
	}); err != nil {
		return Bookmark{}, fmt.Errorf("failed to save bookmark %w", err)
	}

	return Bookmark{
		Title:       title,
		Url:         u,
		Description: description,
	}, nil
}

func (m *BookmarkManager) ListBookmarks(ctx context.Context) ([]Bookmark, error) {
	rows, err := m.repo.GetBookmarks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get bookmarks: %w", err)
	}

	var bookmarks []Bookmark
	for _, row := range rows {
		u, err := url.Parse(row.Url)
		if err != nil {
			return nil, fmt.Errorf("failed to parse url: %w", err)
		}
		bookmarks = append(bookmarks, Bookmark{
			Title:       row.Title,
			Url:         u,
			Description: row.Description.String,
		})
	}

	return bookmarks, nil
}

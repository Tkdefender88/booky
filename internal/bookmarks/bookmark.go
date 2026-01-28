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
	ID          int64
	Title       string
	Url         *url.URL
	Description string
}

func (m *BookmarkManager) SaveBookmark(
	ctx context.Context,
	title, uri, description string,
	tags []string,
) (Bookmark, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return Bookmark{}, fmt.Errorf("failed to parse url: %w", err)
	}

	bookmarkID, err := m.repo.CreateBookmark(ctx, generated.CreateBookmarkParams{
		Title:       title,
		Url:         u.String(),
		Description: sql.NullString{String: description, Valid: true},
	})
	if err != nil {
		return Bookmark{}, fmt.Errorf("failed to save bookmark %w", err)
	}

	if err := m.saveTags(ctx, tags, bookmarkID); err != nil {
		return Bookmark{
			ID:          bookmarkID,
			Title:       title,
			Url:         u,
			Description: description,
		}, fmt.Errorf("failed to save tags: %w", err)
	}

	return Bookmark{
		ID:          bookmarkID,
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
			ID:          row.ID,
			Title:       row.Title,
			Url:         u,
			Description: row.Description.String,
		})
	}

	return bookmarks, nil
}

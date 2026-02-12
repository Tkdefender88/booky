package bookmarks

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Tkdefender88/booky/internal/repo/generated"
)

type Tag struct {
	ID   int64
	Name string
}

func (b *BookmarkManager) saveTags(ctx context.Context, tags []string, bookmarkID int64) error {
	var err error
	tags = append(tags, "all")
	for _, tag := range tags {
		tagID, createErr := b.repo.CreateTag(ctx, strings.ToLower(tag))
		if createErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to save tag %q: %w", tag, createErr))
		}

		if insertErr := b.repo.InsertBookmarkTagJunction(ctx, generated.InsertBookmarkTagJunctionParams{
			BookmarkID: bookmarkID,
			TagID:      tagID,
		}); insertErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to associate bookmark '%d' with tag %q: %w", bookmarkID, tag, insertErr))
		}
	}
	return err
}

func (b *BookmarkManager) ListTags(ctx context.Context) ([]string, error) {
	return b.repo.GetTags(ctx)
}

package bookmarks

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Tkdefender88/booky/internal/repo/generated"
)

func (b *BookmarkManager) saveTags(ctx context.Context, tags []string, bookmarkID int64) error {
	var err error
	tags = append(tags, "all")
	for _, tag := range tags {
		if err := b.repo.CreateTag(ctx, strings.ToLower(tag)); err != nil {
			errors.Join(err, fmt.Errorf("failed to save tag %q: %w", tag, err))
		}

		if err := b.repo.InsertBookmarkTagJunction(ctx, generated.InsertBookmarkTagJunctionParams{
			BookmarkID: bookmarkID,
			Tag:        tag,
		}); err != nil {
			errors.Join(err, fmt.Errorf("failed to associate bookmark '%d' with tag %q: %w", bookmarkID, tag, err))
		}
	}
	return err
}

func (b *BookmarkManager) ListTags(ctx context.Context) ([]string, error) {
	return b.repo.GetTags(ctx)
}

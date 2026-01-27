package bookmarks

import (
	"context"
	"errors"
	"fmt"
)

func (b *BookmarkManager) saveTags(ctx context.Context, tags []string) error {
	var err error
	for _, tag := range tags {
		if err := b.repo.CreateTag(ctx, tag); err != nil {
			errors.Join(err, fmt.Errorf("failed to save tag %q: %w", tag, err))
		}
	}
	return err
}

func (b *BookmarkManager) ListTags(ctx context.Context) ([]string, error) {
	return b.repo.GetTags(ctx)
}

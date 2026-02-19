-- +goose Up
-- +goose StatementBegin
ALTER TABLE tags ADD COLUMN tag_order INTEGER NOT NULL DEFAULT 999;

UPDATE tags SET tag_order = 0 WHERE tag_name = 'all';

UPDATE tags SET tag_order = (
  SELECT COUNT(*) FROM tags t2 
  WHERE t2.id <= tags.id AND t2.tag_name != 'all'
) WHERE tag_name != 'all';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tags DROP COLUMN tag_order;
-- +goose StatementEnd

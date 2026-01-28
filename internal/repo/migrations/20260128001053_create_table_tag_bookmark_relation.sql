-- +goose Up
-- +goose StatementBegin
create table bookmarks_tags (
  tag_name text not null references tag(name),
  bookmark_id integer not null references bookmark(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bookmarks_tags;
-- +goose StatementEnd

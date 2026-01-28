-- +goose Up
-- +goose StatementBegin
create table bookmarks (
  id integer primary key autoincrement,
  title text not null,
  url text not null,
  description text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bookmarks;
-- +goose StatementEnd

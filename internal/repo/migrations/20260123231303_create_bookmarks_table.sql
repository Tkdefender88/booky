-- +goose Up
-- +goose StatementBegin
create table bookmarks (
  id int primary key,
  title text not null,
  url text not null,
  description text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bookmarks;
-- +goose StatementEnd

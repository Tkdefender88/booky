-- +goose Up
-- +goose StatementBegin
create table tags (
  id integer primary key,
  tag_name text not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tags;
-- +goose StatementEnd

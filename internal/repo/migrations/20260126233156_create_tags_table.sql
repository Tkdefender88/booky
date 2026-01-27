-- +goose Up
-- +goose StatementBegin
create table tags (
  tag_name text not null primary key
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tags;
-- +goose StatementEnd

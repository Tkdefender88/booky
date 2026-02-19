-- +goose Up
-- +goose StatementBegin
create table bookmarks_tags (
  tag_id integer not null,
  bookmark_id integer not null,

  primary key (tag_id, bookmark_id),

  constraint fk_bookmarks_tags_bookmarks
    foreign key (bookmark_id) references bookmarks(id)
    on delete cascade,

  constraint fk_bookmarks_tags_tags
    foreign key (tag_id) references tags(id)
    on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bookmarks_tags;
-- +goose StatementEnd

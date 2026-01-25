-- name: GetBookmarks :many
select title, url, description from bookmarks order by id desc;

-- name: CreateBookmark :exec
insert into bookmarks (title, url, description) values (:title, :url, :description);

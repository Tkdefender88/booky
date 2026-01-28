-- name: GetBookmarks :many
select id, title, url, description from bookmarks order by id desc;

-- name: CreateBookmark :one
insert into bookmarks (title, url, description) values (:title, :url, :description) returning id;

-- name: InsertBookmarkTagJunction :exec
insert into bookmarks_tags (bookmark_id, tag_name) values (:bookmark_id, :tag);

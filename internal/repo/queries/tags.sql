-- name: GetTags :many
select tag_name from tags order by tag_name;

-- name: CreateTag :one
insert into tags (tag_name) values (:tag_name) on conflict(tag_name) do update set tag_name=excluded.tag_name returning id;

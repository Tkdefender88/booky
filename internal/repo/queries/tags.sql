-- name: GetTags :many
select tag_name from tags order by tag_name;

-- name: CreateTag :exec
insert or ignore into tags (tag_name) values (:tag_name);

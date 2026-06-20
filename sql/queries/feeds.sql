-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;


-- name: GetFeeds :many
SELECT f.name, f.url, u.name AS "created_by"
FROM feeds AS f
JOIN
  users AS u
  ON u.id = f.user_id;


-- name: GetFeed :one
SELECT id
FROM feeds
WHERE url = $1;


-- name: MarkFeedFetched :exec
UPDATE feeds
SET
  last_fetched_at = $2,
  updated_at = $2
WHERE id = $1;

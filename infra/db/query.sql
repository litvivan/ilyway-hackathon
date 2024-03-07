-- name: GetItem :one
SELECT * FROM items
WHERE id = $1;

-- name: ListItems :many
SELECT * FROM items
WHERE
    participant_count >= coalesce(sqlc.narg('min_participant_count')::int, participant_count) AND
    participant_count <= coalesce(sqlc.narg('max_participant_count')::int, participant_count) AND
    activity_type = coalesce(sqlc.narg('activity_type')::text, activity_type) AND
    author_rating >= coalesce(sqlc.narg('min_author_rating')::float, author_rating) AND
    city = coalesce(sqlc.narg('city')::text, city) AND
    start_at >= coalesce(sqlc.narg('min_start_at')::timestamp, start_at) AND
    start_at <= coalesce(sqlc.narg('max_start_at')::timestamp, start_at)
ORDER BY id;

-- name: InsertItem :one
INSERT INTO items (title, description, participant_count, activity_type, city, author_name, author_rating, image_url, full_address, has_reward, duration, start_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;
-- name: GetItem :one
SELECT * FROM items
WHERE id = $1;

-- name: ListItems :many
SELECT * FROM items
ORDER BY id;

-- name: InsertItem :one
INSERT INTO items (title, description, participant_count, activity_type, city, author_name, author_rating, image_url, full_address, has_reward, duration, start_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;
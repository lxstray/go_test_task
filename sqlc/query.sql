-- name: SelectTopBanner :one
SELECT * FROM banners
WHERE geo = $1 AND feature = $2
ORDER BY cpm DESC
LIMIT 1;

-- name: SelectAll :many
SELECT * FROM banners;

-- name: SelectById :one
SELECT * FROM banners
WHERE id = $1
LIMIT 1;

-- name: CreateBanner :one
INSERT INTO banners (name, image, cpm, geo, feature)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateBanner :one
UPDATE banners
SET name = $2, image = $3, cpm = $4, geo = $5, feature = $6
WHERE id = $1
RETURNING *;

-- name: DeleteBanner :exec
DELETE FROM banners
WHERE id = $1;
-- name: GetTreesByLocation :many
SELECT *
FROM dataTree
WHERE latitude BETWEEN $1 AND $2
  AND longitude BETWEEN $3 AND $4;

-- name: ListTreesFromDatabase :many
SELECT * FROM dataTree
ORDER BY created_at DESC;
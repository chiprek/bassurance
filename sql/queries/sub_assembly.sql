-- name: CreateSubAssembly :one
INSERT INTO sub_assemblies (
id, unit_id, name, serial_number, status
)
VALUES ( $1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSubAssemblies :many
SELECT * FROM sub_assemblies
WHERE unit_id = $1 AND deleted_at IS NULL;

-- name: CreateSubAssembly :one
INSERT INTO sub_assemblies (
id, unit_id, name, serial_number, status
)
VALUES ( $1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateProject :one
INSERT INTO project (id, tenant_id, name, description, status, created_by, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: GetProjectByIDAndTenantID :one
SELECT * FROM project
WHERE id = $1
AND tenant_id = $2;

-- name: ListProjectsByTenantID :many
SELECT * FROM project
WHERE tenant_id = $1
ORDER BY created_at DESC;

-- name: UpdateProject :one
UPDATE project
SET name = $1, description = $2, status = $3, updated_at = $4
WHERE id = $5
AND tenant_id = $6
RETURNING *;

-- name: DeleteProjectByIDAndTenantID :exec
DELETE FROM project
WHERE id = $1
AND tenant_id = $2;
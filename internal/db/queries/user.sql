-- name: CreateUser :one
INSERT INTO users(id, email, is_sso, full_name, password_hash, status, last_login_at, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: CheckIfUserExistsByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: FindUserOneByEmail :one
SELECT * FROM users where email = $1;


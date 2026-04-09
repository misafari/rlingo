-- name: GetLanguageByCode :one
SELECT * FROM language where code = $1;

-- name: GetAllLanguages :many
SELECT * FROM language;

-- name: CheckIfLanguageExistsByCode :one
SELECT EXISTS(SELECT 1 FROM language WHERE code = $1);

-- name: CreateLanguage :one
INSERT INTO language(code, name, native_name, rtl)
VALUES ($1, $2, $3, $4)
RETURNING code;

-- name: UpdateLanguage :one
UPDATE language
SET name = $1, native_name = $2, rtl = $3
WHERE code = $4
RETURNING code;
-- name: CreateUser :one
INSERT INTO users(username, password, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT id,
    email,
    username,
    password
FROM users
WHERE email = $1;
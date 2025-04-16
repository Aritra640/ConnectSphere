-- name: AddUser :one
INSERT INTO users (username , email , password_hashed)
VALUES ($1 , $2 , $3)
RETURNING id, username, email, password_hashed, created_at;

-- name: GetUserByUsername :one 
SELECT id, username, email, password_hashed, created_at
FROM users 
WHERE username = $1;

-- name: GetUserbyEmail :one 
SELECT id, username, email, password_hashed, created_at
FROM users 
WHERE email = $1;

-- name: GetUserByID :one 
SELECT id, username, email, password_hashed, created_at 
FROM users 
WHERE ID = $1;

-- name: GetUsersAll :many 
SELECT id, username, email, password_hashed, created_at
FROM users 
ORDER BY created_at desc;

-- name: DeleteUserByID :exec 
DELETE FROM users WHERE id = $1;

-- name: UpdateUserPasswordByUsername :exec 
UPDATE users 
  set password_hashed = $2
WHERE username = $1;

-- name: UpdateUserUsernameByEmail :exec 
UPDATE users 
  set username = $2
WHERE email = $1;


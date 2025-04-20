-- name: CreateNewRefreshToken :one 
INSERT INTO refresh_token (user_id , token , expires_at)
VALUES ($1 , $2 , $3)
RETURNING *;


-- name: GetRefreshToken :one
SELECT * FROM refresh_token WHERE token = $1;


-- name: DeleteRefreshTokenByToken :exec 
DELETE FROM refresh_token WHERE token = $1;


-- name: DeleteRefreshTokenByUserID :exec 
DELETE FROM refresh_token WHERE user_id = $1;


-- name: GetRefreshTokenByUserId :one 
SELECT * FROM refresh_token WHERE user_id = $1;

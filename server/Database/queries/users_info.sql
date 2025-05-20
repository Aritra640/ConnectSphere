-- name: AddUserInfo :one
INSERT INTO users_info (user_id , pimage, pbio)
VALUES ($1 , $2 , $3) 
RETURNING *;


-- name: GetUserInfo :one 
SELECT * FROM users_info 
WHERE user_id = $1;

-- name: UpdateUserInfo :exec 
UPDATE users_info 
  set pimage = $2,
      pbio = $3
WHERE user_id = $1;

-- name: UpdateUserInfoImage :exec 
UPDATE users_info 
  set pimage = $2
WHERE user_id = $1;

-- name: UpdateUserInfoBio :exec 
UPDATE users_info 
  set pbio = $2
WHERE user_id = $1;

-- name: DeleteUserInfo :exec 
DELETE FROM users_info WHERE user_id = $1;


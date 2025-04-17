-- name: AddUserToGroup :exec 
INSERT INTO chat_group_member (group_id, user_id , is_admin)
VALUES ($1 , $2 , $3)
ON CONFLICT (group_id , user_id) DO NOTHING;

-- name: RemoveUserFromGroup :exec 
DELETE FROM chat_group_member WHERE user_id = $1 AND group_id = $2;

-- name: GetGroupMembers :many 
SELECT u.*
FROM chat_group_member cm
JOIN users u ON cm.user_id = u.id
WHERE cm.group_id = $1;

-- name: GetAppAdmins :many 
SELECT u.*
FROM chat_group_member cm
JOIN users u ON cm.user_id = u.id
WHERE cm.group_id = $1 AND cm.is_admin = TRUE;

-- name: IsUserInGroup :one 
SELECT EXISTS (
  SELECT 1 FROM chat_group_member
  WHERE user_id = $1 AND group_id = $2
) AS is_member;


-- name: SetUserToAdminStatus :exec
UPDATE chat_group_member
SET is_admin = $3
WHERE user_id = $1 AND group_id = $2;


-- name: IsUserGroupAdmin :one
SELECT EXISTS (
  SELECT 1 FROM chat_group_member
  WHERE user_id = $1 AND group_id = $2 AND is_admin = TRUE
) AS is_admin;


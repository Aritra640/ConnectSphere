-- name: AddGroupJoinRequest :exec 
INSERT INTO chat_group_queue (user_id, chat_group_id) 
VALUES ($1, $2);

-- name: AcceptGroupJoinRequest :exec 
UPDATE chat_group_queue
SET is_accepted = TRUE
WHERE user_id = $1 AND chat_group_id = $2;

-- name: DeleteChatGroupRequest :exec 
DELETE FROM chat_group_queue
WHERE user_id = $1 AND chat_group_id = $2;

-- name: GetPendingChatGroupRequestsForGroup :many 
SELECT * FROM chat_group_queue
WHERE chat_group_id = $1 AND is_accepted = FALSE;

-- name: GetGroupRequest :one 
SELECT * FROM chat_group_queue
WHERE user_id = $1 AND chat_group_id = $2; 

-- name: ListUsersAllGroupRequest :many 
SELECT * FROM chat_group_queue
WHERE user_id = $1;

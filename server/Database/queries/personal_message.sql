-- name: CreatePersonalMessage :one
INSERT INTO personal_message (chat_id , sender_id , receiver_id)
VALUES ($1 , $2 , $3) 
RETURNING *;

-- name: GetMessagesBetweenTwoUsers :many 
SELECT c.*, pm.is_seen, pm.send_at
FROM personal_message pm
JOIN chat c ON c.id = pm.chat_id
WHERE (pm.sender_id = $1 AND pm.receiver_id = $2)
   OR (pm.sender_id = $2 AND pm.receiver_id = $1)
ORDER BY pm.send_at;

-- name: MarkMessageAsSeen :exec
UPDATE personal_message
SET is_seen = TRUE
WHERE chat_id = $1; 

-- name: DeletePersonalMessage :exec 
DELETE FROM personal_message
WHERE chat_id = $1;


-- name: EditMessageContent :exec 
UPDATE chat
SET content = $2
WHERE id = $1;

-- name: GetUnseenMessage :many 
SELECT c.*, pm.send_at
FROM personal_message pm
JOIN chat c ON c.id = pm.chat_id
WHERE pm.receiver_id = $1 AND pm.is_seen = FALSE
ORDER BY pm.send_at;

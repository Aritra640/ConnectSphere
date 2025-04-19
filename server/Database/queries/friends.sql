-- name: GetUserFriends :many
SELECT u.*
FROM users u
JOIN friends f ON u.id = f.friend_id
WHERE f.user_id = $1;

-- name: AddUserFriend :exec 
INSERT INTO friends (user_id, friend_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: AdduserFriendsBothWays :exec 
INSERT INTO friends (user_id, friend_id)
VALUES 
  ($1, $2),
  ($2, $1)
ON CONFLICT DO NOTHING;

-- name: AreFriends :one
SELECT EXISTS (
  SELECT 1 FROM friends
  WHERE user_id = $1 AND friend_id = $2
) AS are_friends;

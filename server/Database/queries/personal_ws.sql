-- name: CreatePersonalWS :one
INSERT INTO personal_ws (id, usera, userb)
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetPersonalWSbyUsers :one 
SELECT * FROM personal_ws 
WHERE (usera = $1 AND userb = $2) OR (usera = $2 AND userb = $1);


-- name: DeletePersonalWSbyID :exec 
DELETE FROM personal_ws WHERE id = $1;


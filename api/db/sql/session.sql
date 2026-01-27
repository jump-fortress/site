-- name: InsertSession :one
insert into session (player_id, token_id)
  values (?, ?)
  returning *;
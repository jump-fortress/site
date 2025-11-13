-- name: AddSession :one
insert into session (player_id, token_id)
  select player.id, ? from player
  where steam_id64 = ?
  returning *;
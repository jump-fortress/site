-- name: InsertPlayer :one
insert into player (steam_id64)
  values (?)
  on conflict do update
  set steam_id64 = steam_id64
  returning *;

-- name: SelectPlayer :one
select * from player
  where id = ?;

-- name: SelectAllPlayers :many
select * from player;

-- name: SelectPlayerFromSteamID64 :one
select * from player
  where steam_id64 = ?;

-- name: UpdatePlayerSessionInfo :one
update player
  set steam_avatar_url = ?,
  display_name = ?
  where steam_id64 = ?
  returning *;

-- name: UpdatePlayerPreferredClassFromSteamID64 :exec
update player
  set preferred_class = ?
  where steam_id64 = ?;
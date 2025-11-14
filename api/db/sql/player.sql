-- name: InsertPlayer :one
insert into player (steam_id64)
  values (?)
  returning *;

-- name: SelectPlayer :one
select * from player
  where id = ?;

-- name: SelectPlayerFromSteamID64 :one
select * from player
  where steam_id64 = ?;

-- name: UpdatePlayerSessionInfo :one
update player
  set steam_avatar_url = ?,
  display_name = ?
  where steam_id64 = ?
  returning *;
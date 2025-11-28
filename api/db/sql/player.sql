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

-- name: UpdatePlayerSteamAvatarURLFromSteamID64 :exec
update player
  set steam_avatar_url = ?
  where steam_id64 = ?;

-- name: UpdatePlayerSteamTradeTokenFromSteamID64 :exec
update player
  set steam_trade_token = ?
  where steam_id64 = ?;

-- name: UpdatePlayerTempusIDFromSteamID64 :exec
update player
  set tempus_id = ?
  where steam_id64 = ?;

-- name: UpdatePlayerDisplayName :exec
update player
  set display_name = ?
  where id = ?;

-- name: UpdatePlayerSoldierDivision :exec
update player
  set soldier_division = ?
  where id = ?;

-- name: UpdatePlayerDemoDivision :exec
update player
  set demo_division = ?
  where id = ?;

-- name: UpdatePlayerPreferredClassFromSteamID64 :exec
update player
  set preferred_class = ?
  where steam_id64 = ?;

-- name: UpdatePlayerPreferredLauncherFromSteamID64 :exec
update player
  set preferred_launcher = ?
  where steam_id64 = ?;
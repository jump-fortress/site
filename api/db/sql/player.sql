-- name: InsertPlayer :one
insert into player (id)
  values (?)
  on conflict do update
    set id = id
  returning *;

-- name: SelectPlayer :one
select * from player
  where id = ?;

-- name: SelectAllPlayers :many
select * from player;

-- name: UpdatePlayerSessionInfo :one
update player
  set steam_avatar_url = ?,
  display_name = ?
  where id = ?
  returning *;

-- name: UpdatePlayerSteamAvatarURL :exec
update player
  set steam_avatar_url = ?
  where id = ?;

-- name: UpdatePlayerSteamTradeToken :exec
update player
  set steam_trade_token = ?
  where id = ?;

-- name: UpdatePlayerTempusInfo :exec
update player
  set tempus_id = ?,
  country = ?,
  country_code = ?
  where id = ?;

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

-- name: UpdatePlayerMotwTimeslot :exec
update player
  set motw_timeslot = ?
  where id = ?;

-- name: UpdatePlayerPreferredClass :exec
update player
  set preferred_class = ?
  where id = ?;

-- name: UpdatePlayerPreferredLauncher :exec
update player
  set preferred_launcher = ?
  where id = ?;

-- name: UpdatePlayerPreferredMap :exec
update player
  set preferred_map = ?
  where id = ?;

-- name: UpdatePlayerRole :exec
update player
  set role = ?
  where id = ?;
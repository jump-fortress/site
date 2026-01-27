-- name: InsertPlayer :one
insert into player (id)
  values (?)
  on conflict do update
    set id = id
  returning *;

-- name: SelectPlayer :one
select * from player
  where id = ?;

-- name: SelectPlayers :many
select * from player;

-- name: UpdatePlayerAlias :exec
update player
  set alias = ?
  where id = ?;

-- name: UpdatePlayerAvatarURL :exec
update player
  set avatar_url = ?
  where id = ?;

-- name: UpdatePlayerSoldierDiv :exec
update player
  set soldier_div = ?
  where id = ?;

-- name: UpdatePlayerDemoDiv :exec
update player
  set demo_div = ?
  where id = ?;

-- name: UpdatePlayerTradeToken :exec
update player
  set trade_token = ?
  where id = ?;

-- name: UpdatePlayerTempusInfo :exec
update player
  set tempus_id = ?,
    country = ?,
    country_code = ?
  where id = ?;

-- name: UpdatePlayerClassPref :exec
update player
  set class_pref = ?
  where id = ?;

-- name: UpdatePlayerMapPref :exec
update player
  set map_pref = ?
  where id = ?;

-- name: UpdatePlayerLauncherPref :exec
update player
  set launcher_pref = ?
  where id = ?;

-- name: UpdatePlayerRole :exec
update player
  set role = ?
  where id = ?;
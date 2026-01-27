-- name: InsertLeaderboard :exec
insert into leaderboard (event_id, div, map)
values (?, ?, ?);

-- name: UpdateLeaderboard :exec
update leaderboard
  set div = ?,
  map = ?
  where id = ?;

-- name: SelectLeaderboard :one
select * from leaderboard
  where id = ?;

-- name: SelectLeaderboards :many
select * from leaderboard
  where event_id = ?;
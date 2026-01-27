-- name: InsertTime :exec
insert into time (leaderboard_id, player_id, tempus_time_id, duration, verified)
  values (?, ?, ?, ?, ?);

-- name: SelectTime :one
select * from time
  where id = ?;

-- name: CountPlayerTimesFromLeaderboard :one
select count(*) from time
  where leaderboard_id = ?
  and player_id = ?;

-- name: SelectPRTime :one
select cast(min(duration) as float) as duration from time
  where leaderboard_id = ?
  and player_id = ?;

-- name: SelectTimeExists :one
select exists(
  select 1 from time
  where leaderboard_id = ?
  and player_id = ?
  and duration = ?
);

-- name: VerifyTime :exec
update time
  set verified = true
  where id = ?;

-- name: DeleteTime :exec
delete from time
  where id = ?;

-- name: SelectTimesFromLeaderboard :many
select * from time
where leaderboard_id = ?;
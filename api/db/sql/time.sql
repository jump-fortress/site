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
select sqlc.embed(time), sqlc.embed(player) from time
  join player on time.player_id = player.id
  where time.leaderboard_id = ?
  order by time.duration asc;

-- name: SelectPRTimesFromLeaderboard :many
select sqlc.embed(time), sqlc.embed(player), cast(rank() over (order by duration) as integer) time_position from time
  join player on time.player_id = player.id
  where time.leaderboard_id = ?
  group by player.id;

-- name: UpdateTimeFromTempus :exec
update time
  set duration = ?,
  tempus_time_id = ?,
  created_at = ?
  where id = ?;

-- name: SelectParticipatedEvents :many
select sqlc.embed(event), sqlc.embed(leaderboard), sqlc.embed(time) from time
join leaderboard on time.leaderboard_id = leaderboard.id
join event on leaderboard.event_id = event.id
where time.player_id = ?
group by leaderboard.id
order by event.starts_at desc;
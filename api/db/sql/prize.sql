-- name: InsertPrize :exec
insert into prize (leaderboard_id, position, keys)
  values (?, ?, ?);

-- name: SelectPrize :one
select * from prize
  where leaderboard_id = ?
  and player_id = ?;

-- name: UpdatePrize :exec
update prize
  set player_id = ?
  where leaderboard_id = ?
  and position = ?;

-- name: DeletePrizepool :exec
delete from prize
  where leaderboard_id = ?;

-- name: SelectPrizepool :many
select * from prize
  where leaderboard_id = ?;

-- name: SelectPrizepoolTotal :one
select cast(coalesce(sum(p.keys), 0) as integer) from prize p
  join leaderboard l on p.leaderboard_id = l.id
  join event e on l.event_id = e.id
  where e.id = ?;
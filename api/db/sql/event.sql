-- name: SelectEvent :one
select * from event
  where id = ?;

-- name: SelectEventFromLeaderboardID :one
select e.* from event e
  join leaderboard l on l.event_id = e.id
  where l.id = ?;

-- name: SelectEventKinds :many
select * from event
  where kind = ?;

-- name: SelectEventKind :one
select * from event
  where kind = ?
  and kind_id = ?;

-- name: SelectEventLeaderboards :many
select sqlc.embed(event), sqlc.embed(leaderboard) from event
  join leaderboard on leaderboard.event_id = event.id
  where event.kind = ?
    and event.kind_id = ?;

-- name: CountEventKinds :one
select count(*) from event
  where kind = ?;

-- name: InsertEvent :exec
insert into event (kind, kind_id, class, visible_at, starts_at, ends_at)
  values (?, ?, ?, ?, ?, ?);

-- name: DeleteEvent :exec
delete from event
  where id = ?;
-- name: SelectEvent :one
select * from event
  where id = ?;

-- name: SelectEvents :many
select * from event;

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

-- name: SelectLastMOTW :one
select * from event
  where kind = 'motw'
  order by starts_at asc;

-- name: SelectEventLeaderboards :many
select sqlc.embed(event), sqlc.embed(leaderboard) from event
  join leaderboard on leaderboard.event_id = event.id
  where event.kind = ?
    and event.kind_id = ?;

-- name: CountEventKinds :one
select count(*) from event
  where kind = ?;

-- name: InsertEvent :one
insert into event (kind, kind_id, class, visible_at, starts_at, ends_at)
  values (?, ?, ?, ?, ?, ?)
  returning *;

-- name: UpdateEvent :exec
update event
  set kind = ?,
  kind_id = ?,
  class = ?,
  visible_at = ?,
  starts_at = ?,
  ends_at = ?
  where id = ?;

-- name: DeleteEvent :exec
delete from event
  where id = ?;
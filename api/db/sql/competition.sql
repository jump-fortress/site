-- name: InsertCompetition :one
insert into competition (class, starts_at, ends_at, visible_at)
  values (?, ?, ?, ?)
  returning *;

-- name: CancelCompetition :one
delete from competition
  where id = ? and starts_at > current_timestamp
  returning *;
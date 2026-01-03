-- name: InsertCompetition :one
insert into competition (class, starts_at, ends_at)
  values (?, ?, ?)
  returning *;
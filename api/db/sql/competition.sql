-- name: InsertCompetition :one
insert into competition (class, starts_at, ends_at, visible_at)
  values (?, ?, ?, ?)
  returning *;

-- name: UpdateCompetition :exec
update competition
  set class = ?,
  starts_at = ?,
  ends_at = ?,
  visible_at = ?
  where id = ? and starts_at < current_timestamp;

-- name: UpdateCompetitionPrizepool :exec
update competition
  set prizepool = ?
  where id = ?;

-- name: UpdateCompetitionComplete :exec
update competition
  set complete = true
  where id = ?;

-- name: DeleteCompetition :one
delete from competition
  where id = ? and starts_at > current_timestamp
  returning *;
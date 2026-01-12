-- name: InsertCompetition :one
insert into competition (type, class, starts_at, ends_at, visible_at)
  values (?, ?, ?, ?, ?)
  returning *;

-- name: SelectCompetition :one
select * from competition
  where id = ?;

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

-- name: CompleteCompetition :exec
update competition
  set complete = true
  where id = ?;

-- name: DeleteCompetition :one
delete from competition
  where id = ? and starts_at > current_timestamp
  returning *;
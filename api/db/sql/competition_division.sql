-- name: InsertCompetitionDivision :exec
insert into competition_division (competition_id, division, map)
  values (?, ?, ?);

-- name: SelectCompetitionDivision :one
select * from competition_division
  where id = ?;

-- name: SelectCompetitionDivisions :many
select * from competition_division
  where competition_id = ?;

-- name: UpdateCompetitionDivision :exec
update competition_division
  set map = ?
  where id = ?;

-- name: DeleteCompetitionDivision :one
delete from competition_division
  where id = ?
  returning *;
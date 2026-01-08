-- name: InsertCompetitionDivision :exec
insert into competition_division (competition_id, division, map)
  values (?, ?, ?);

-- name: UpdateCompetitionDivision :exec
update competition_division
  set map = ?
  where competition_id = ? and division = ?;

-- name: DeleteCompetitionDivision :exec
delete from competition_division
  where competition_id = ? and division = ?;

-- name: SelectCompetitionDivisions :many
select * from competition_division
  where competition_id = ?;
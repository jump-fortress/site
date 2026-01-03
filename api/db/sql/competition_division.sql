-- name: InsertCompetitionDivision :exec
insert into competition_division (competition_id, division, map)
values (?, ?, ?);

-- name: SelectCompetitionDivisions :many
select * from competition_division
  where competition_id = ?;
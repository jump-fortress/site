-- name: InsertCompetitionPrize :exec
insert into competition_prize (competition_division_id, placement, amount)
  values (?, ?, ?);

-- name: SelectCompetitionPrizes :many
select * from competition_prize
  where competition_division_id = ?;

-- name: UpdateCompetitionPrize :exec
update competition_prize
  set placement = ?,
  amount = ?
  where competition_division_id = ?;

-- name: DeleteCompetitionPrize :exec
delete from competition_prize
  where competition_division_id = ?;
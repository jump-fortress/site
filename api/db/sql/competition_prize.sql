-- name: InsertCompetitionPrize :exec
insert into competition_prize (competition_division_id, placement, amount)
  values (?, ?, ?);

-- name: SelectDivisionPrizepool :many
select * from competition_prize cp
  join competition_division cd on cp.competition_division_id = cd.id
  where cd.id = ?;

-- name: UpdateCompetitionPrize :exec
update competition_prize
  set placement = ?,
  amount = ?
  where competition_division_id = ?;

-- name: DeleteCompetitionDivisionPrizepool :exec
delete from competition_prize
  where competition_division_id = ?;

-- name: DeleteCompetitionPrize :exec
delete from competition_prize
  where id = ?;

-- name: SumCompetitionPrizepool :one
select cast(coalesce(sum(amount),0) as integer) from competition_prize cp
  join competition_division cd on cp.competition_division_id = cd.id
  join competition c on cd.competition_id = c.id
  where c.id = ?;
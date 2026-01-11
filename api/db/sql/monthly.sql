-- name: InsertMonthly :exec
insert into monthly (competition_id)
  values (?);

-- name: SelectMonthly :one
select * from monthly m
  join competition c on m.competition_id = c.id
  where m.id = ?;

-- name: SelectAllMonthly :many
select * from monthly m
  join competition c on m.competition_id = c.id;
-- name: InsertMap :exec
insert or ignore into map (id, name, courses, bonuses, soldier_tier, demo_tier, soldier_rating, demo_rating)
  values (?, ?, ?, ?, ?, ?, ?, ?);

-- name: SelectMaps :many
select * from map;

-- name: SelectMapNames :many
select name from map;
-- name: InsertMap :exec
insert or ignore into map (id, name, courses, bonuses, soldier_tier, demo_tier, soldier_rating, demo_rating)
  values (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetMaps :many
select * from map;
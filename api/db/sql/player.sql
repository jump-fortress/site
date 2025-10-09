-- name: InsertPlayer :one
insert or ignore into player (steam_id, display_name, soldier_division, demo_division)
  values (?, ?, ?, ?)
  returning *;
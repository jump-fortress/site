-- name: InsertPlayer :one
insert or ignore into player (steam_id, steam_pfp_id, display_name)
  values (?, ?, ?)
  returning *;

-- name: SelectPlayer :one
select * from player
  where id = ?;
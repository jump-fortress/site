-- name: InsertPlayer :one
insert into player (steam_id64)
  values (?)
  on conflict do update set steam_id64 = steam_id64
  returning *;

-- name: SelectPlayer :one
select * from player
  where id = ?;
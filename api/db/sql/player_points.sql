-- name: InsertPlayerPoints :exec
insert or ignore into player_points (class, player_id)
values (?, ?);
  
-- name: SelectPlayerPoints :one
select * from player_points
where player_id = ?
and class = ?;
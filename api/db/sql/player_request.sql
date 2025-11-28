-- name: InsertPlayerRequest :exec
insert into player_request (player_id, type, content)
  select player.id, ?, ? from player
  where player.steam_id64 = ?;

-- name: GetAllPendingPlayerRequests :many
select * from player_request
  where id = ? and pending = true;

-- name: GetAllPlayerRequests :many
select * from player_request;
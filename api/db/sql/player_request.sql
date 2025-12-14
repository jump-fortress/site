-- name: InsertPlayerRequest :exec
insert into player_request (player_id, type, content)
values (?, ?, ?);

-- may be unnecessary
-- name: GetPendingPlayerRequestsForPlayer :many
select * from player_request
where player_id = ? and pending = true;

-- name: CheckPendingPlayerRequest :one
select exists (
  select 1
  from player_request
  where player_id = ? and type = ? and pending = true);

-- name: GetAllPlayerRequests :many
select * from player_request;
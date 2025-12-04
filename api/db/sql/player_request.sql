-- name: InsertPlayerRequest :exec
insert into player_request (player_id, type, content)
values (?, ?, ?);

-- name: GetAllPendingPlayerRequests :many
select * from player_request
where player_id = ? and pending = true;

-- name: GetAllPlayerRequests :many
select * from player_request;
-- name: InsertPlayerRequest :exec
insert into player_request (player_id, type, content)
values (?, ?, ?);

-- may be unnecessary
-- name: SelectPendingPlayerRequestsForPlayer :many
select * from player_request
where player_id = ? and pending = true;

-- name: CheckPendingPlayerRequest :one
select exists (
  select 1
  from player_request
  where player_id = ? and type = ? and pending = true);

-- name: SelectAllPendingPlayerRequests :many
select * from player_request pr
  join player p on pr.player_id = p.id
  where pr.pending = true;

-- name: ResolvePlayerRequest :exec
update player_request
  set pending = false
  where id = ?;
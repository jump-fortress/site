-- name: InsertRequest :exec
insert into request (player_id, kind, content)
  values (?, ?, ?);

-- name: SelectPlayerRequests :many
select * from request
  where player_id = ?;

-- name: CheckPendingRequestExists :one
select exists (
  select 1
  from request
  where player_id = ? and kind = ? and pending = true);

-- name: SelectPendingRequests :many
select * from request
  where pending = true;

-- name: ResolveRequest :exec
update request
  set pending = false
  where id = ?;
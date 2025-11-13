-- name: DisallowToken :exec
insert into disallow_token (token_id)
  values (?);

-- name: GetDisallowToken :one
select exists(
  select 1 from disallow_token
    where token_id = ?
);
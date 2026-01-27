-- name: InsertDisallowToken :exec
insert into disallow_token (token_id)
  values (?);

-- name: SelectDisallowToken :one
select exists(
  select 1 from disallow_token
    where token_id = ?
);
-- name: InsertAuditLog :exec
insert into audit_log(from_player_id, to_player_id, kind, from_content, to_content)
  values(?,?,?,?,?);

-- name: SelectAuditLogs :many
select sqlc.embed(p1), sqlc.embed(p2), sqlc.embed(l) from audit_log l
join player p1 on p1.id = l.from_player_id
join player p2 on p2.id = l.to_player_id
order by l.created_at desc;
-- name: SelectPlayerTimeslot :one
select sqlc.embed(motw_timeslot), sqlc.embed(player_motw_timeslot) from player_motw_timeslot
  join motw_timeslot on player_motw_timeslot.timeslot_id = motw_timeslot.id
where player_motw_timeslot.player_id = ?;

-- name: SelectTimeslots :many
select * from motw_timeslot;

-- name: SelectFirstTimeslot :one
select * from motw_timeslot
order by id asc;

-- name: SelectLastTimeslot :one
select * from motw_timeslot
order by id desc;

-- name: UpdatePlayerTimeslot :exec
update player_motw_timeslot
  set timeslot_id = ?
  where player_id = ?;

-- name: UpsertTimeslot :exec
insert into motw_timeslot (id, starts_at)
  values (?, ?)
  on conflict do update
    set starts_at = excluded.starts_at;
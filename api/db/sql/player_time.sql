-- name: InsertPlayerTime :exec
insert into player_time (player_id, competition_division_id, run_time, verified)
  values (?, ?, ?, ?);

-- name: VerifyPlayerTime :exec
update player_time
  set verified = true
  where id = ?;

-- name: SelectCompetitionDivisionPRTimes :many
select pt.id, pt.player_id, pt.competition_division_id, min(run_time) as run_time, pt.verified, pt.created_at,
p.steam_avatar_url, p.display_name
from player_time pt
  join player p on pt.player_id = p.id
    where pt.competition_division_id = ?
  group by p.id
  order by run_time;

-- name: DeletePlayerTime :one
delete from player_time
  where id = ?
  returning *;
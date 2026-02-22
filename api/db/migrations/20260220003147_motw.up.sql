create table motw_timeslot(
  id integer not null primary key autoincrement,
  starts_at datetime not null
);

create table player_motw_timeslot(
  timeslot_id integer not null default 1,
  player_id text not null unique,

  foreign key (timeslot_id) references motw_timeslot (id),
  foreign key (player_id) references player (id)
);
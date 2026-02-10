create table prize(
  leaderboard_id integer not null,
  player_id integer not null,
  position integer not null,
  amount float not null,
  keys boolean not null,

  foreign key (leaderboard_id) references leaderboard (id),
  foreign key (player_id) references player (id)
);

create table bounty(
  event_id integer not null primary key,
  course integer,
  bonus integer,
  target_duration float,
  complete boolean not null default false,

  created_at datetime not null default current_timestamp,

  foreign key (event_id) references event (id)
);
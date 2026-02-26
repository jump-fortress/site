create table prize(
  leaderboard_id integer not null,
  player_id text,
  position integer not null,
  keys integer not null,

  foreign key (leaderboard_id) references leaderboard (id),
  foreign key (player_id) references player (id)
);
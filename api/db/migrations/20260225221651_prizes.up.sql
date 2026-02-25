create table prize(
  leaderboard_id integer not null,
  player_id integer not null,
  position integer not null,
  amount float not null,
  keys boolean not null,

  foreign key (leaderboard_id) references leaderboard (id),
  foreign key (player_id) references player (id)
);
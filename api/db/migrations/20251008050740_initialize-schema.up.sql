-- initial schema
-- player
create table player(
  id integer not null primary key autoincrement,
  steam_id text unique not null,
  steam_trade_token text unique,
  tempus_id integer unique,
  discord_id text unique,
  display_name text not null,
  soldier_division text,
  demo_division text,
  created_at datetime not null default current_timestamp
);
-- competition
create table competition(
  id integer not null primary key autoincrement,
  class text not null,
  starts_at datetime not null,
  ends_at datetime not null,
  created_at datetime not null default current_timestamp,
  check (class in ("soldier", "demo"))
);
-- "stardust points" for a player
-- relates to player (id)
create table player_points(
  id integer not null primary key autoincrement,
  player_id integer not null,
  soldier_points integer not null default 0,
  demo_points not null default 0,
  foreign key (player_id) references player (id)
);
-- time for a competition's player
-- relates to player (id) and competition (id)
create table player_time(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  player_id integer not null,
  run_time float not null,
  verified boolean not null,
  created_at datetime not null default current_timestamp
);
-- info for a division's competition
-- relates to competition (id)
create table competition_division(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  division text not null,
  map text not null,
  foreign key (competition_id) references competition (id)
);
-- end result for a competition (placement and points)
-- relates to competition (id) and player (id)
create table competition_result(
  competition_id integer not null,
  player_id integer not null,
  placement integer not null,
  points integer not null,
  created_at datetime not null default current_timestamp,
  primary key (competition_id, player_id),
  foreign key (competition_id) references competition (id),
  foreign key (player_id) references player (id)
);
-- prize for a competition (keys)
-- relates to competition (id)
create table competition_prize(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  placement integer not null,
  amount integer not null,
  foreign key (competition_id) references competition (id)
);
-- marathon, a type of competition
-- relates to competition (id)
create table marathon(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  foreign key (competition_id) references competition (id)
);
-- map of the week, a type of competition
-- relates to competition (id)
create table motw(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  foreign key (competition_id) references competition (id)
);
-- bounty, a type of competition with no division relation
-- relates to competition (id)
create table bounty(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  map text not null,
  type text not null,
  time float,
  foreign key (competition_id) references competition (id),
  check (type in ("target time", "record"))
);
-- quest, a type of competition
-- relates to competition (id)
create table quest(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  type text not null,
  time float,
  completion_limit,
  foreign key (competition_id) references competition (id),
  check (type in ("target time", "completion"))
);
-- deleted record from any table
create table deleted_record(
  id integer not null primary key autoincrement,
  source_table text not null,
  source_id text not null,
  data jsonb not null,
  deleted_at datetime not null default current_timestamp
);
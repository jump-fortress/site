-- initial schema

-- player
-- todo: consider what "roles" should exist, and if this is the correct implementation
create table player(
  id integer not null primary key autoincrement,
  role text not null default 'player',
  steam_id64 text unique not null,
  -- steam_avatar_url can be non-unique for default and points shop avatars
  steam_avatar_url text,
  steam_trade_token text unique,
  tempus_id integer unique,
  discord_id text unique,
  display_name text,
  soldier_division text,
  demo_division text,
  preferred_class text not null default 'Soldier',

  created_at datetime not null default current_timestamp

  check (preferred_class in ('Soldier', 'Demo'))
);

-- competition
create table competition(
  id integer not null primary key autoincrement,
  class text not null,
  starts_at datetime not null,
  ends_at datetime not null,

  created_at datetime not null default current_timestamp,

  check (class in ('Soldier', 'Demo'))
);

-- "stardust points" for a player
-- relates to player (id)
create table player_points(
  id integer not null primary key autoincrement,
  class text not null,
  player_id integer not null,
  total integer not null default 0,
  last_9_motw integer not null default 0,
  last_3_monthly integer not null default 0,

  foreign key (player_id) references player (id),

  check (class in ('Soldier', 'Demo'))
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

-- monthly, a type of competition
-- relates to competition (id)
create table monthly(
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
  check (type in ('Target Time', 'Record'))
);

-- quest, a type of competition
-- relates to competition (id)
create table quest(
  id integer not null primary key autoincrement,
  competition_id integer not null,
  type text not null,
  time float,
  completion_limit text not null,

  foreign key (competition_id) references competition (id),
  check (type in ('Target Time', 'Completion'))
);

-- deleted record from any table
create table deleted_record(
  id integer not null primary key autoincrement,
  source_table text not null,
  source_id text not null,
  data jsonb not null,

  deleted_at datetime not null default current_timestamp
);

-- openid_nonce
create table openid_nonce(
  id integer not null primary key autoincrement,
  endpoint text not null,
  nonce_time datetime not null,
  nonce_string text not null,

  created_at datetime not null default current_timestamp,
  
  unique (endpoint, nonce_string)
);

-- player openid session
create table session(
  id integer not null primary key autoincrement,
  player_id integer not null,
  token_id text not null unique,
  
  created_at datetime not null default current_timestamp,
  foreign key (player_id) references player (id)
);

-- player token blacklist
create table disallow_token(
  token_id text not null unique,

  created_at datetime not null default current_timestamp,

  foreign key (token_id) references session (token_id)
);

-- triggers
-- todo: view instead?

-- total points
create trigger update_player_points_total_after_insert 
after insert on competition_result
begin
  update player_points
  set total = (
    select sum(points) from competition_result cr
    join competition c on cr.competition_id = c.id
    where cr.player_id = new.player_id
      and c.class = player_points.class
  )
  where player_id = new.player_id;
end;

-- last 9 motw
create trigger update_player_points_last_9_motw_after_insert
after insert on competition_result
begin
  update player_points
  set last_9_motw = (
    select sum(points) from competition_result cr
    join competition c on cr.competition_id = c.id
    join motw m on m.competition_id = c.id
    where cr.player_id = new.player_id
      and c.class = player_points.class
    order by cr.created_at desc
    limit 9
  )
  where player_id = new.player_id;
end;

-- last 3 monthly
create trigger update_player_points_last_3_monthly_after_insert
after insert on competition_result
begin
  update player_points
  set last_3_monthly = (
    select sum(points) from competition_result cr
    join competition c on cr.competition_id = c.id
    join monthly m on m.competition_id = c.id
    where cr.player_id = new.player_id
      and c.class = player_points.class
    order by cr.created_at desc
    limit 3
  )
  where player_id = new.player_id;
end;

-- todo: deleted record table?


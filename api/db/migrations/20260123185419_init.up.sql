create table player(
  id text not null primary key,
  role text not null default 'player',
  alias text,
  soldier_div text,
  demo_div text,
  avatar_url text,
  trade_token text unique,
  tempus_id integer unique,
  country text,
  country_code text,
  class_pref text not null default 'soldier',
  map_pref text,
  launcher_pref text,

  created_at datetime not null default current_timestamp,

  check (role in ('player', 'mod', 'dev', 'admin')),
  check (class_pref in ('soldier', 'demo'))
);

create table event(
  id integer not null primary key autoincrement,
  kind text not null,
  kind_id integer not null,
  class text not null,

  visible_at datetime not null,
  starts_at datetime not null,
  ends_at datetime not null,

  created_at datetime not null default current_timestamp,

  check (class in ('soldier', 'demo')),
  unique (kind, kind_id)
);

create table leaderboard(
  id integer not null primary key,
  event_id integer not null,
  div text,
  map text not null,

  foreign key (event_id) references event (id) on delete cascade
);

create table time(
  id integer not null primary key autoincrement,
  leaderboard_id integer not null,
  player_id text not null,
  tempus_time_id integer,
  duration float not null,
  verified boolean not null,

  created_at datetime not null default current_timestamp,
  
  foreign key (leaderboard_id) references leaderboard (id),
  foreign key (player_id) references player (id)
);

create table map(
  name text not null unique,
  courses integer,
  bonuses integer, 
  soldier_tier integer not null,
  demo_tier integer not null,
  soldier_rating integer not null,
  demo_rating integer not null
);

create table deleted_row(
  id integer not null primary key autoincrement,
  source_table text not null,
  source_id text not null,
  data jsonb not null,

  deleted_at datetime not null default current_timestamp
);

create table openid_nonce(
  id integer not null primary key autoincrement,
  endpoint text not null,
  nonce_time datetime not null,
  nonce_string text not null,

  created_at datetime not null default current_timestamp,
  
  unique (endpoint, nonce_string)
);

create table session(
  id integer not null primary key autoincrement,
  player_id text not null,
  token_id text not null unique,
  
  created_at datetime not null default current_timestamp,
  foreign key (player_id) references player (id)
);

create table disallow_token(
  token_id text not null unique,

  created_at datetime not null default current_timestamp,

  foreign key (token_id) references session (token_id)
);
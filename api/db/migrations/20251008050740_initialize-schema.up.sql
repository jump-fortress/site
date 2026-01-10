-- initial schema

-- # player info
-- id: steam_id64
-- motw_timeslot: preferred timeslot for motws, in motw_times
create table player(
  id text not null primary key,
  role text not null default 'Player',
  steam_avatar_url text,
  steam_trade_token text unique,
  tempus_id integer unique,
  country text,
  country_code text,
  discord_id text unique,
  display_name text,
  soldier_division text,
  demo_division text,
  motw_timeslot integer,
  preferred_class text not null default 'Soldier',
  preferred_launcher text not null default 'None',
  preferred_map text,

  created_at datetime not null default current_timestamp,

  check (role in ('Player', 'Consultant', 'Moderator', 'Treasurer', 'Admin')),
  check (motw_timeslot in (1, 2, 3)),
  check (preferred_class in ('Soldier', 'Demo')),
  check (preferred_launcher in ('Stock', 'Original', 'Mangler', 'None'))
);

-- # competition info that applies to each kind 
-- (monthly, motw, bounty, quest, ..., are not autoincremented)
-- this is so they may follow a # order even when deleted (canceled)
--
-- shouldn't be deleted after starts_at
--
-- visible_at: visible to non-admins, without competition_division maps listed
-- complete: manually set to true after ends_at, generate competition_result
-- prizepool: total prizepool (keys), updated when competition_prizepool is updated
create table competition(
  id integer not null primary key autoincrement,
  class text not null,
  starts_at datetime not null,
  ends_at datetime not null,
  visible_at datetime not null,
  complete boolean not null default false,
  prizepool integer,

  created_at datetime not null default current_timestamp,

  check (class in ('Soldier', 'Demo'))
);

-- # maps pulled from Tempus data
-- courses: may be null (a map with no courses is basically one course, so course PRs don't matter)
-- bonuses: may be null (a map with no bonuses.. has no bonuses)
create table map(
  id integer not null primary key,
  name text not null,
  courses integer,
  bonuses integer,
  soldier_tier integer not null,
  demo_tier integer not null,
  soldier_rating integer not null,
  demo_rating integer not null
);

-- not implemented
-- # special achievement earned on profiles 
-- (ex. world cup winner)
--
-- label: tooltip text
-- href: badge link, if appropriate
create table badge(
  id integer not null primary key,
  label text not null,
  href text,
  src text not null,

  created_at datetime not null default current_timestamp
);

-- not implemented
-- # player badge
create table player_badge(
  id integer not null primary key,
  player_id text not null,
  badge_id integer not null,
  achieved_at datetime not null,

  created_at datetime not null default current_timestamp,

  foreign key (player_id) references player (id),
  foreign key (badge_id) references badge (id)
);

-- # "stardust points" for a player
-- achieved through some competitions (monthly, motw)
-- updated when a competition is "complete"
create table player_points(
  id integer not null primary key autoincrement,
  player_id text not null,
  class text not null,
  total integer not null default 0,
  last_9_motw integer not null default 0,
  last_3_monthly integer not null default 0,

  foreign key (player_id) references player (id),

  check (class in ('Soldier', 'Demo')),
  unique (class, player_id)
);

-- # player request
create table player_request(
  id integer not null primary key autoincrement,
  player_id text not null,
  type text not null,
  content text,
  pending boolean not null default true,

  created_at datetime not null default current_timestamp,

  foreign key (player_id) references player (id),
  check (type in ('Display Name Change', 'Soldier Placement', 'Demo Placement'))
);

-- # time for a competition's player
-- not a final result, but a table of all times submitted
create table player_time(
  id integer not null primary key autoincrement,
  player_id text not null,
  competition_division_id integer not null,
  run_time float not null,
  verified boolean not null,

  created_at datetime not null default current_timestamp,

  foreign key (player_id) references player (id),
  foreign key (competition_division_id) references competition_division (id)
);

-- # division and map pair for a competition (excluding bounty)
create table competition_division(
  id integer not null primary key,
  competition_id integer not null,
  division text not null,
  map text not null,

  foreign key (competition_id) references competition (id) on delete cascade
);

-- not implemented
-- # final result for a competition
--
-- prize_id: can be null, not every competition or placement has a prize
create table competition_result(
  time_id integer not null,
  prize_id integer,
  points integer,

  created_at datetime not null default current_timestamp,

  primary key (time_id, prize_id),
  foreign key (time_id) references player_time (id),
  foreign key (prize_id) references competition_prize (id)
);

-- # prize for a competition (keys)
create table competition_prize(
  id integer not null primary key,
  competition_division_id integer not null,
  placement integer not null,
  amount integer not null,

  foreign key (competition_division_id) references competition_division (id) on delete cascade
);

-- # monthly, a type of competition
create table monthly(
  id integer not null primary key,
  competition_id integer not null,

  foreign key (competition_id) references competition (id) on delete cascade
);

-- not implemented
-- # map of the week, a type of competition
create table motw(
  id integer not null primary key,
  competition_id integer not null,

  foreign key (competition_id) references competition (id) on delete cascade
);

-- not implemented
-- # timeslots limit when a motw is available to view and submit times for
create table motw_timeslot(
  id integer not null primary key autoincrement,
  starts_at datetime not null,
  ends_at datetime not null
);

-- not implemented
-- # bounty, a type of competition with no division relation
-- if course and bonus are null, bounty is for a map run
create table bounty(
  id integer not null primary key,
  competition_id integer not null,
  map text not null,
  course integer,
  bonus integer,
  type text not null,
  time float,
  
  foreign key (competition_id) references competition (id) on delete cascade,
  check (type in ('Target Time', 'Record'))
);

-- not implemented
-- # quest, a type of competition
create table quest(
  id integer not null primary key,
  competition_id integer not null,
  type text not null,
  time float,
  completion_limit text not null,

  foreign key (competition_id) references competition (id) on delete cascade,
  check (type in ('Target Time', 'Completion'))
);

-- # deleted record from any table
create table deleted_record(
  id integer not null primary key autoincrement,
  source_table text not null,
  source_id text not null,
  data jsonb not null,

  deleted_at datetime not null default current_timestamp
);

-- # openid_nonce
create table openid_nonce(
  id integer not null primary key autoincrement,
  endpoint text not null,
  nonce_time datetime not null,
  nonce_string text not null,

  created_at datetime not null default current_timestamp,
  
  unique (endpoint, nonce_string)
);

-- # player openid session
create table session(
  id integer not null primary key autoincrement,
  player_id text not null,
  token_id text not null unique,
  
  created_at datetime not null default current_timestamp,
  foreign key (player_id) references player (id)
);

-- # player token blacklist
create table disallow_token(
  token_id text not null unique,

  created_at datetime not null default current_timestamp,

  foreign key (token_id) references session (token_id)
);


-- # triggers

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
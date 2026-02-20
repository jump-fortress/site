create table request(
  id integer not null primary key autoincrement,
  player_id text not null,
  kind text not null,
  content text not null,
  pending boolean not null default true,

  created_at datetime not null default current_timestamp,

  foreign key (player_id) references player (id),
  check (kind in ('alias update', 'soldier div', 'demo div'))
);
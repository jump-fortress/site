create table requests(
  id integer not null primary key autoincrement,
  player_id text not null,
  type text not null,
  content text,
  pending boolean not null default true,

  created_at datetime not null default current_timestamp,

  foreign key (player_id) references player (id),
  check (type in ('alias update', 'soldier div', 'demo div'))
);
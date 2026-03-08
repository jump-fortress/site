create table audit_log(
  from_player_id text not null,
  to_player_id text not null,
  kind text not null,
  from_content text not null,
  to_content text not null,

  created_at datetime not null default current_timestamp,

  foreign key (from_player_id) references player (id),
  foreign key (to_player_id) references player (id)
);
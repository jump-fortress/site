create table points(
  time_id integer not null,
  prize_id integer,
  points float not null,
  position integer not null,

  foreign key (time_id) references time (id),
  foreign key (prize_id) references prize (id)
);
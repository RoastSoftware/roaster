set search_path to roaster;

do $$
begin
  -- Almost everything with an @ is a valid e-mail, so what the hell.
  create domain email as text check (value ~ '@' and
				     char_length(value) < 255 and
				     char_length(value) > 2); --

  create domain username as text check (value !~* E'[\\s:;,+$\\/\\\\?!*\'()@=&#]' and
					char_length(value) > 0 and
					char_length(value) <= 30); --

  create domain fullname as text check(char_length(value) < 255); --

  create domain code as text check(char_length(value) <= 500000); --

  create domain score as integer check (value >= 0); --

  exception when others then
    raise notice 'domains already exists, skipping...'; --
end
$$;

create table if not exists "user"
(
  username username not null
    constraint user_pkey
    primary key,
  hash     bytea not null,
  create_time	timestamp with time zone not null,
  fullname fullname,
  email    email not null unique
);

create unique index if not exists username_user_idx on "user" (lower(username));

create table if not exists user_followees
(
    username username not null,
    create_time timestamp with time zone not null,
    followee username not null,
    constraint followee_relation_uq unique (username, followee),
    constraint username_fk foreign key (username)
        references "user" (username) match simple
        on update cascade
        on delete cascade,
    constraint followee_fk foreign key (username)
        references "user" (username) match simple
        on update cascade
        on delete cascade
);

create index if not exists username_user_followees_idx
on user_followees using btree (username);

create table if not exists "roast"
(
  id       	serial        	not null
    constraint roast_pk
    primary key,
  code     	code 		not null,
  username 	username	not null
    constraint user_fk
    references "user" (username)
    on update cascade
    on delete cascade,
  score    	score 		not null,
  language 	text          	not null,
  create_time	timestamp	with time zone not null
);

create table if not exists "warning"
(
  id		uuid    not null
    constraint warning_pk
    primary key,
  row         	integer not null,
  "column"    	integer not null,
  engine      	text    not null,
  name        	text    not null,
  description 	text    not null
);

create table if not exists "error"
(
  id		uuid    not null
    constraint error_pk
    primary key,
  row         	integer not null,
  "column"    	integer not null,
  engine      	text    not null,
  name        	text    not null,
  description text    not null
);

create table if not exists "roast_statistics"
(
  roast integer not null
    constraint roast_fk
    references roast (id)
    on update cascade
    on delete cascade,
  lines_of_code integer not null,
  number_of_errors integer not null,
  number_of_warnings integer not null
);

create table if not exists "roast_has_errors"
(
  roast integer not null
    constraint roast_fk
    references roast (id)
    on update cascade
    on delete cascade,
  error uuid not null
    constraint error_fk
    references error (id)
);

create table if not exists "roast_has_warnings"
(
  roast   integer not null
    constraint roast_fk
    references roast (id)
    on update cascade
    on delete cascade,
  warning uuid not null
    constraint warning_fk
    references warning (id)
);

create table if not exists avatar
(
  avatar bytea not null,
  username username not null,
  constraint username_uq unique (username),
  constraint username_fk foreign key (username)
    references "user" (username) match simple
    on update cascade
    on delete cascade
);

create index if not exists username_avatar_idx
  on avatar using btree (username);

drop function if exists round_minutes(timestamp without time zone, integer);
drop function if exists round_minutes(timestamp without time zone, integer, text);

create function round_minutes(timestamp without time zone, integer)
returns timestamp without time zone as $$
  select
     date_trunc('hour', $1)
     +  cast(($2::varchar||' min') as interval)
     * round(
     (date_part('minute',$1)::float + date_part('second',$1)/ 60.)::float
     / $2::float
      )
$$ language sql immutable;

create function round_minutes(timestamp without time zone, integer, text)
returns text as $$
  select to_char(round_minutes($1,$2),$3)
$$ language sql immutable;

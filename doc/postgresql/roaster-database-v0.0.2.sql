set search_path to roaster;

create table if not exists "user"
(
  username text  not null
    constraint user_pkey
    primary key
    constraint username_chk
    check (char_length(username) <= 30),
  hash     bytea not null,
  create_time	timestamp with time zone not null,
  fullname text
    constraint fullname_chk
    check (char_length(fullname) < 255),
  email    text  not null unique
    constraint email_chk
    check (char_length(email) < 255)
);

create unique index if not exists username_user_idx on "user" (lower(username));

create table if not exists user_friends
(
    username text not null,
    create_time timestamp with time zone not null,
    constraint username_fk foreign key (username)
        references "user" (username) match simple
        on update cascade
        on delete cascade
);

create index if not exists username_user_friends_idx
on user_friends using btree (username);

create table if not exists "roast"
(
  id       	serial        	not null
    constraint roast_pk
    primary key,
  code     	text          	not null
    constraint code_chk
    check (char_length(code) <= 500000),
  username 	text          	not null
    constraint user_fk
    references "user" (username)
    on update cascade
    on delete cascade,
  score    	integer 	not null
    constraint score_chk
    check (score >= 0),
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
  username text not null,
  constraint username_uq unique (username),
  constraint username_fk foreign key (username)
    references "user" (username) match simple
    on update cascade
    on delete cascade
);

create index if not exists username_avatar_idx
  on avatar using btree (username);

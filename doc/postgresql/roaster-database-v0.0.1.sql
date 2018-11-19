create database roaster
with owner postgres;

create sequence roaster.roast_id_seq
  as integer
  maxvalue 2147483647;

alter sequence roaster.roast_id_seq
  owner to postgres;

create sequence roaster.error_id_seq
  as integer
  maxvalue 2147483647;

alter sequence roaster.error_id_seq
  owner to postgres;

create sequence roaster.warning_id_seq
  as integer
  maxvalue 2147483647;

alter sequence roaster.warning_id_seq
  owner to postgres;

create table if not exists roaster."user"
(
  username text  not null
    constraint user_pkey
    primary key
    constraint username_chk
    check (char_length(username) <= 30),
  hash     bytea not null,
  fullname text
    constraint fullname_chk
    check (char_length(fullname) < 255),
  email    text  not null
    constraint email_chk
    check (char_length(email) < 255)
);

alter table roaster."user"
  owner to postgres;

create table if not exists roaster.roast
(
  id       serial        not null
    constraint roast_pk
    primary key,
  code     text          not null
    constraint code_chk
    check (char_length(code) <= 500000),
  username text          not null
    constraint user_fk
    references "user",
  score    numeric(5, 4) not null
    constraint score_chk
    check ((score >= (0) :: numeric) AND (score <= (1) :: numeric)),
  language text          not null
);

alter table roaster.roast
  owner to postgres;

create table if not exists roaster.warning
(
  hash        bytea   not null,
  row         integer not null,
  "column"    integer not null,
  engine      text    not null,
  name        text    not null,
  description text    not null,
  id          serial  not null
    constraint warning_pk
    primary key
);

alter table roaster.warning
  owner to postgres;

create unique index if not exists warning_hash_idx
  on roaster.warning (hash);

create table if not exists roaster.error
(
  hash        bytea   not null,
  row         integer not null,
  "column"    integer not null,
  engine      text    not null,
  name        text    not null,
  description text    not null,
  id          serial  not null
    constraint error_pk
    primary key
);

alter table roaster.error
  owner to postgres;

create unique index if not exists error_hash_idx
  on roaster.error (hash);

create table if not exists roaster.roast_has_errors
(
  roast integer not null
    constraint roast_fk
    references roast (id),
  error integer not null
    constraint error_fk
    references error
);

alter table roaster.roast_has_errors
  owner to postgres;

create table if not exists roaster.roast_has_warnings
(
  roast   integer not null
    constraint roast_fk
    references roast (id),
  warning integer not null
    constraint warning_fk
    references warning
);

alter table roaster.roast_has_warnings
  owner to postgres;



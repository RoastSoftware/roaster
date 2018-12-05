// Package forwardengineer was generated automatically by inlinesql at 2018-12-05 13:39:51.6108381 +0100 CET m=+0.012690801.
package forwardengineer
// GetQueries returns a pre-parsed slice of SQL queries.
func GetQueries() []string {
	return []string{
		"set search_path to roaster",
		"create sequence if not exists roast_id_seq as integer maxvalue 2147483647",
		"create sequence if not exists error_id_seq as integer maxvalue 2147483647",
		"create sequence if not exists warning_id_seq as integer maxvalue 2147483647",
		"create table if not exists \"user\" ( username text not null constraint user_pkey primary key constraint username_chk check (char_length(username) <= 30), hash bytea not null, create_time timestamp with time zone not null, fullname text constraint fullname_chk check (char_length(fullname) < 255), email text not null constraint email_chk check (char_length(email) < 255) )",
		"create table if not exists roast ( id serial not null constraint roast_pk primary key, code text not null constraint code_chk check (char_length(code) <= 500000), username text not null constraint user_fk references \"user\" (username), score numeric(5, 4) not null constraint score_chk check ((score >= (0) :: numeric) AND (score <= (1) :: numeric)), language text not null )",
		"create table if not exists warning ( hash bytea not null, row integer not null, \"column\" integer not null, engine text not null, name text not null, description text not null, id serial not null constraint warning_pk primary key )",
		"create unique index if not exists warning_hash_idx on warning (hash)",
		"create table if not exists error ( hash bytea not null, row integer not null, \"column\" integer not null, engine text not null, name text not null, description text not null, id serial not null constraint error_pk primary key )",
		"create unique index if not exists error_hash_idx on error (hash)",
		"create table if not exists roast_has_errors ( roast integer not null constraint roast_fk references roast (id), error integer not null constraint error_fk references error (id) )",
		"create table if not exists roast_has_warnings ( roast integer not null constraint roast_fk references roast (id), warning integer not null constraint warning_fk references warning (id) )",
		"create table if not exists avatar ( avatar bytea not null, username text not null, constraint username_uq unique (username), constraint username_fk foreign key (username) references \"user\" (username) match simple on update no action on delete no action )",
		"create index if not exists username_idx on avatar using btree (username)",
	}
}

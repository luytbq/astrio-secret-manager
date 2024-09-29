create table keys (
    id serial primary key,
    value varchar(256),
    create_at timestamptz
);

alter table secrets
add column description varchar(256),
add column decrypt boolean not null default false,
add column key_id integer;
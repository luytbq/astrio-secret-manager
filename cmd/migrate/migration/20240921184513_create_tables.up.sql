create table secret_groups (
    id serial primary key,
    user_id int,
    description varchar(2048)    
);

create table secrets (
    id serial primary key,
    group_id integer references secret_groups(id),
    content varchar(2048)
);
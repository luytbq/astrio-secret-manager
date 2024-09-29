drop table keys;

alter table secrets
drop column description,
drop column decrypt,
drop column key_id;
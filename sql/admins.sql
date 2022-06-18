drop table if exists admins;
create table admins(
    full_name varchar(30),
    email varchar(50) primary key,
    password_hash varchar(512),
    access_level int default 0
);
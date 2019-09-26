create table users (
                       id uuid primary key,
                       login varchar(50) not null unique,
                       password varchar(50) not null
);
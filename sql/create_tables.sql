CREATE TABLE profile (
                       id uuid PRIMARY KEY,
                       name varchar(50)
);

create table users (
                       id uuid primary key,
                       login varchar(50) not null unique,
                       password varchar(50) not null
);
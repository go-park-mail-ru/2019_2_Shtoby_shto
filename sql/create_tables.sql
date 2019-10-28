create table users (
                       id uuid primary key,
                       login varchar(50) not null unique,
                       password varchar(50) not null,
                       photo_id uuid
);

create table photo (
                       id uuid primary key,
                       time_load timestamp with time zone,
                       path varchar(50)
);

create table board (
                       id uuid primary key,
                       caption varchar(50)
);
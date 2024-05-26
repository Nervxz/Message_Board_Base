/*
DB migration on service startup:
--------------------------------
- Django, Ruby on Rails, Spring Boot: only recommended for development.
- Hard to write a good migration script.
- In case of multi instances/distributed system, on new deployment, which server
instance will handle the migration?
*/

drop schema if exists msg_board cascade;

create schema msg_board;

create table if not exists msg_board.users
(
    id         serial primary key,
    username   varchar(255) unique not null,
    password   bytea               not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

create table if not exists msg_board.topics
(
    id         serial primary key,
    by         integer references msg_board.users (id),
    title      varchar(255) not null,
    body       text         not null,
    created_at timestamp with time zone
);

create table if not exists msg_board.comments
(
    id         serial primary key,
    by         integer references msg_board.users (id),
    content    text not null,
    topic_id   integer references msg_board.topics (id),
    created_at timestamp with time zone
);

create table if not exists msg_board.votes
(
    id         serial primary key,
    by         integer references msg_board.users (id),
    topic_id   integer references msg_board.topics (id),
    unique (by, topic_id),
    created_at timestamp with time zone
);

-- +goose Up
create table if not exists users (
    id serial primary key,
    name text not null,
    email text unique not null,
    password text not null,
    role text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- +goose Down
drop table if exists users;
-- +goose Up
-- +goose StatementBegin
create extension if not exists citext;

create table users (
    id bigint primary key generated by default as identity,
    email citext not null unique,
    password text not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd

-- +goose Up

IF NOT EXISTS create table users (
    id          bigint      primary key,
    firstname    varchar(64),
    created_at  timestamp
);

-- +goose Down

drop table users;
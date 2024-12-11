-- +goose Up
-- +goose StatementBegin
create table if not exists access_endpoints (
    id serial primary key,
    role varchar(16) not null check (role <> ''),
    endpoint_address varchar(64) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists access_endpoints;
-- +goose StatementEnd

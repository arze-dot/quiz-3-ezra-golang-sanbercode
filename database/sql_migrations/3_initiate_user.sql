
-- +migrate Up
-- +migrate StatementBegin

create table users (
    id           BIGINT NOT NULL,
    username     varchar(256),
    password     varchar(256),
    created_at   TIMESTAMP,
    created_by   varchar(256),
    modified_at  TIMESTAMP,
    modified_by  varchar(256)
)

-- +migrate StatementEnd

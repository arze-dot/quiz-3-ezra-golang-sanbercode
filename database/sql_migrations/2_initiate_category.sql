
-- +migrate Up
-- +migrate StatementBegin

create table category (
    id           BIGINT NOT NULL,
    name         varchar(256),
    created_at   TIMESTAMP,
    created_by   varchar(256),
    modified_at  TIMESTAMP,
    modified_by  varchar(256)
)

-- +migrate StatementEnd

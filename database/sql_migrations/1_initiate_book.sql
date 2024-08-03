
-- +migrate Up
-- +migrate StatementBegin

create table book (
    id            BIGINT NOT NULL,
    title         varchar(256),
    description   varchar(256),
    image_url     varchar(256),
    release_year  INT,
    price         INT,
    total_page    INT,
    thickness     INT,
    category_id   BIGINT,
    created_at    TIMESTAMP,
    created_by    varchar(256),
    modified_at   TIMESTAMP,
    modified_by   varchar(256)
)

-- +migrate StatementEnd

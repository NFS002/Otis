-- groups table schema
CREATE TABLE IF NOT EXISTS groups (
        id                     SERIAL,
        name                   VARCHAR         not null,
        description            VARCHAR,
        created_at             DATE            not null       default current_date,
        deleted_at             DATE

        CHECK       ( name <> '' ),
        CONSTRAINT  pk_groups_id   PRIMARY KEY (id)
);
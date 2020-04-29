-- groupmembers table schema
CREATE TABLE IF NOT EXISTS groupmembers (
        group_id               INTEGER         not null,
        user_id                INTEGER         not null,
        created_at             DATE            not null       default current_date,
        deleted_at             DATE,

        CONSTRAINT dulplicate_members UNIQUE (group_id, user_id),

        CONSTRAINT  pk_groupmembers_id         PRIMARY KEY (group_id),
        CONSTRAINT  fk_groumembers_id          FOREIGN KEY (group_id)  REFERENCES groups (id) ON UPDATE CASCADE ON DELETE RESTRICT,
        CONSTRAINT  fk_groumembers_user_id     FOREIGN KEY (user_id)   REFERENCES users (id)  ON UPDATE CASCADE ON DELETE RESTRICT
);
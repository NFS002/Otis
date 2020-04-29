
-- Install pgcrypto extension https://blog.dbi-services.com/securely-store-passwords-in-postgresql/
CREATE EXTENSION pgcrypto;


-- accounts table schema
CREATE TABLE IF NOT EXISTS accounts (
        user_id            INTEGER           not null        UNIQUE,
        username           VARCHAR           not null,
        password           VARCHAR           not null,
        email              VARCHAR           not null        UNIQUE,
        created_at         DATE              not null        default current_date,
        deleted_at         DATE

        CHECK ( length(username) > 5 ),
        CHECK ( length(password) > 10 ),
        CHECK ( length(email) <> 5 ),

        CONSTRAINT         pk_accounts_user_id     PRIMARY KEY(user_id),
        CONSTRAINT         fk_accounts_user_id     FOREIGN KEY(user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE RESTRICT
)








        graduation_year    INT                 not null        default 0,
                                     photo_url               VARCHAR             not null        default '',
                                     first_name              VARCHAR(100)        not null,
                                     last_name               VARCHAR(100)        not null        default '',
                                     alias                   VARCHAR(100)        not null        default 'anon',
                                     nationality             VARCHAR             not null,
                                     average_weekly_spend    DECIMAL             not null        default 0,
                                     expense_bands           VARCHAR             not null        default ''

                                         CHECK (first_name <> ''),
                                     CHECK (alias <> ''),
                                     CHECK (average_weekly_spend > 0),

                                     CONSTRAINT pk_users_id PRIMARY KEY (id),
                                     CONSTRAINT fk_users_university_id       FOREIGN KEY (university_id) REFERENCES universities (id) ON UPDATE CASCADE ON DELETE RESTRICT,
                                     CONSTRAINT fk_users_nationality         FOREIGN KEY (nationality)   REFERENCES countries    (code) ON UPDATE CASCADE ON DELETE RESTRICT
);
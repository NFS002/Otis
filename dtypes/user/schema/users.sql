-- gender enum data type
create type GENDER_T as enum('FEMALE', 'MALE', 'OTHER');

-- users table schema
CREATE TABLE IF NOT EXISTS users (
    id                      SERIAL,
    date_of_birth           DATE                not null,
    gender                  GENDER_T            not null,
    gender_description      VARCHAR(100)        not null        default '',
    university_id           INT                 not null,
    created_at              DATE                not null        default current_date,
    graduation_year         INT                 not null        default 0,
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
    CONSTRAINT fk_users_nationality         FOREIGN KEY (nationality)   REFERENCES countries    (id) ON UPDATE CASCADE ON DELETE RESTRICT
);
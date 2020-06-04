-- gender enum data type
create type MERCHANT_T as enum('PARTNER', 'MERCHANT');

-- merchants table schema
CREATE TABLE IF NOT EXISTS users (
    id                      SERIAL,
    created_at              DATE                not null,
    status                  MERCHANT_T          not null,
    acquired_at             DATE,
    name                    VARCHAR             not null,
    sector                  VARCHAR             not null,
    street_address          VARCHAR             not null,
    country                 VARCHAR             not null         default 'GB',
    postcode                VARCHAR(10)         not null,
    city                    VARCHAR             not null,
    longitude               VARCHAR             not null         default '0',
    latitude                VARCHAR             not null         default '0'

    CHECK ( name <> '')
    CHECK ( sector <> '')
    CHECK ( street_address <> '' ),
    CHECK ( postcode <> '' ),
    CHECK ( city <> '' ),

    CONSTRAINT pk_merchants_id PRIMARY KEY (id)
);
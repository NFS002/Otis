CREATE TABLE IF NOT EXISTS universities (
    id                     SERIAL,
    name                   VARCHAR     not null,
    street_address         VARCHAR     not null,
    postcode               VARCHAR     not null,
    city                   VARCHAR     not null,
    longitude              VARCHAR     not null         default '0',
    latitude               VARCHAR     not null         default '0'

    CHECK ( name <> '' ),

    CONSTRAINT pk_universities_id PRIMARY KEY (id)
);
-- gender enum data type
create type MERCHANT_T as enum('PARTNER', 'MERCHANT');

-- merchants table schema
CREATE TABLE IF NOT EXISTS users (
    id                      SERIAL,
    created_at              DATE                not null,
    status                  MERCHANT_T          not null,
    acquired_at             DATE                not null
);
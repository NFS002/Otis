-- gender enum data type
create type TRANSACTION_T as enum('DEBIT', 'CREDIT');

CREATE TABLE IF NOT EXISTS transactions (
    id                      SERIAL,
    created_at              DATE            not null,
    created_at_time         TIMESTAMP       ,
    processed_time          TIMESTAMP,
    outlet_id               INTEGER,
    value                   DECIMAL,
    currency                VARCHAR(5)      not null,
    is_online               BOOLEAN,
    is_contactless          BOOLEAN,
    type                    TRANSACTION_T   not null,
    offer_id                INTEGER,

    CHECK ( value > 0 )
)
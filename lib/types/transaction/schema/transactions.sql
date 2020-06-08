-- gender enum data type
create type TRANSACTION_T as enum('DEBIT', 'CREDIT', 'UNKNOWN');

CREATE TABLE IF NOT EXISTS transactions (
    id                      SERIAL,
    created_at              DATE            not null       default current_date,
    created_at_time         TIME            not null       default current_time,
    processed_date          DATE,
    processed_time          TIME,
    outlet_id               INTEGER         not null,
    user_id                 INTEGER         not null,
    value                   DECIMAL         not null,
    currency                VARCHAR(5)      not null        default 'GBP',  -- ISO 4217
    is_online               BOOLEAN         not null        default false,
    is_contactless          BOOLEAN         not null        default false,
    type                    TRANSACTION_T   not null        default 'UNKNOWN',

    CHECK ( value > 0 ),

    CONSTRAINT pk_transactions_id           PRIMARY KEY (id),
    CONSTRAINT fk_transactions_outlet_id    FOREIGN KEY (outlet_id) REFERENCES outlets (outlet_id)  ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_transactions_users_id     FOREIGN KEY (user_id)   REFERENCES users   (user_id)    ON UPDATE CASCADE ON DELETE RESTRICT,
)
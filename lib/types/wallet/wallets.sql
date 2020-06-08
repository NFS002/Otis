-- wallets table schema
CREATE TABLE IF NOT EXISTS wallets (
        id                     SERIAL               not null,
        balance                DECIMAL              not null,
        created_at             TIMESTAMP            not null,
        updated_at             TIMESTAMP            not null,

        CONSTRAINT             pk_wallet_id         PRIMARY KEY (id),
        CONSTRAINT             positive_balance     CHECK (balance >= 0)
);

-- drop wallet_created_trigger if exists
DROP TRIGGER IF EXISTS prevent_update_trigger ON wallets RESTRICT;

-- create wallet_created_trigger using wallet_auto_create()
CREATE TRIGGER prevent_update_trigger BEFORE INSERT OR UPDATE ON wallets
    FOR EACH ROW EXECUTE PROCEDURE prevent_update();


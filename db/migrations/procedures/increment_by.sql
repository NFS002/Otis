-- create or replace increment_wallet_by()
CREATE OR REPLACE FUNCTION increment_wallet_by( wallet_id INTEGER, d DECIMAL ) RETURNS void AS
$BODY$
BEGIN
    UPDATE wallets
    SET balance = balance + d
    WHERE id = wallet_id;
END
$BODY$
    LANGUAGE plpgsql;
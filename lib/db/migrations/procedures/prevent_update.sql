-- create or replace prevent_update() function
CREATE OR REPLACE FUNCTION prevent_update()
    RETURNS trigger AS
$BODY$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        IF NEW.created_at IS NOT NULL OR NEW.updated_at IS NOT NULL THEN
            RAISE WARNING 'You may not set the value of the created_at or updated_at columns manually and these values will be overwritten';
        end if;
        NEW.created_at = NOW();
        NEW.updated_at = NOW();
    ELSEIF (TG_OP = 'UPDATE') THEN
        IF NEW.created_at <> OLD.created_at OR NEW.updated_at <> OLD.updated_at THEN
            RAISE WARNING 'You may not set the value of the created_at or updated_at columns manually and these values will be overwritten';
        END IF;
        NEW.created_at = OLD.created_at;
        NEW.updated_at = NOW();
    end if;
    RETURN NEW;
END
$BODY$
    LANGUAGE plpgsql;
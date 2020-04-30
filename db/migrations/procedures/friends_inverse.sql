-- friends inverse function
CREATE OR REPLACE FUNCTION friends_inverse()
    RETURNS trigger AS
$BODY$
BEGIN
    IF EXISTS( SELECT * FROM friends as f WHERE (new.user_id = f.friend_id  AND new.friend_id = f.user_id)) THEN
        RAISE EXCEPTION 'The same (inversed) relationship already exists in this table';
    END IF;
    RETURN NEW;
END
$BODY$
    LANGUAGE plpgsql;
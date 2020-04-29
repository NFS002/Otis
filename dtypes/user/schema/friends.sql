-- friends table schema
CREATE TABLE IF NOT EXISTS friends
(
    user_id    INTEGER not null,
    friend_id  INTEGER not null,
    created_at DATE    not null default CURRENT_DATE,
    deleted_at DATE,


    CONSTRAINT pk_friends_id_id PRIMARY KEY (user_id, friend_id),
    CONSTRAINT              fk_friends_user_id     FOREIGN KEY (user_id)    REFERENCES users (user_id) ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT              fk_friends_friend_id   FOREIGN KEY (friend_id)  REFERENCES users (user_id) ON UPDATE CASCADE ON DELETE RESTRICT
);


-- friends inverse function
CREATE OR REPLACE FUNCTION check_friends_inverse()
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

-- friends inverse trigger
CREATE TRIGGER friends_inverse_trigger BEFORE INSERT OR UPDATE ON friends
    FOR EACH ROW EXECUTE PROCEDURE check_friends_inverse();
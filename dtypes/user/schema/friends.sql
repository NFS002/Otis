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

-- friends inverse trigger
CREATE TRIGGER friends_inverse_trigger BEFORE INSERT OR UPDATE ON friends
    FOR EACH ROW EXECUTE PROCEDURE friends_inverse();
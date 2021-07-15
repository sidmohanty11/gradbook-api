BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS messages (
    id serial PRIMARY KEY,
    from_user_id integer NOT NULL,
    to_user_id integer NOT NULL,
    content TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(from_user_id) REFERENCES users (id) on delete cascade,
    FOREIGN KEY(to_user_id) REFERENCES users (id) on delete cascade
);

COMMIT;

BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS messages (
    id serial PRIMARY KEY,
    from_user_id serial NOT NULL,
    to_user_id serial NOT NULL,
    content TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(from_user_id) REFERENCES users (id),
    FOREIGN KEY(to_user_id) REFERENCES users (id)
);

COMMIT;

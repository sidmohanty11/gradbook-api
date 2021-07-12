BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS questions (
    id serial PRIMARY KEY NOT NULL,
	user_id integer NOT NULL,
    q_text TEXT NOT NULL,
	created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id) on delete cascade
);

COMMIT;

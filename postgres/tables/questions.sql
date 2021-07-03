BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS questions (
    id serial PRIMARY KEY,
	user_id serial NOT NULL,
    q_text TEXT NOT NULL,
	created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

COMMIT;

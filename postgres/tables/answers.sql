BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS answers (
    id serial PRIMARY KEY,
    q_id serial NOT NULL,
	user_id serial NOT NULL,
    a_text TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(q_id) REFERENCES questions(id)
);

COMMIT;
BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS answers (
    id serial PRIMARY KEY,
    q_id integer NOT NULL,
	user_id integer NOT NULL,
    a_text TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) on delete cascade,
    FOREIGN KEY(q_id) REFERENCES questions(id) on delete cascade
);

COMMIT;
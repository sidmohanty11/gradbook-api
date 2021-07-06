BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS blogs (
    id serial PRIMARY KEY,
	user_id integer NOT NULL,
    blog_title TEXT NOT NULL,
    blog_text TEXT NOT NULL,
	created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

COMMIT;

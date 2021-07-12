BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS stories (
    id serial PRIMARY KEY,
	user_id integer NOT NULL,
    name VARCHAR(100) NOT NULL,
    branch VARCHAR(50) NOT NULL,
    clubs VARCHAR(100) NOT NULL,
    motto TEXT NOT NULL,
    github_link TEXT DEFAULT '' NOT NULL,
    youtube_link TEXT DEFAULT '' NOT NULL,
    linkedin_link TEXT DEFAULT '' NOT NULL,
    image_url TEXT NOT NULL,
    journey TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) on delete cascade
);

COMMIT;

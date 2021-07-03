BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS stories (
    id serial PRIMARY KEY,
	user_id serial NOT NULL,
    name VARCHAR(100) NOT NULL,
    branch VARCHAR(50) NOT NULL,
    clubs VARCHAR(100) NOT NULL,
    motto TEXT NOT NULL,
    github_link TEXT,
    youtube_link TEXT,
    linkedin_link TEXT,
    image_url TEXT NOT NULL,
    journey TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

COMMIT;

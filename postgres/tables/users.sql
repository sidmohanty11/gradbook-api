BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
	username VARCHAR (50) UNIQUE NOT NULL,
    hash VARCHAR (200) NOT NULL,
	email VARCHAR (255) UNIQUE NOT NULL,
    image_url VARCHAR (500) NOT NULL,
	created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    last_login TIMESTAMP
);

COMMIT;
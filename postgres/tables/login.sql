BEGIN TRANSACTION;

CREATE TABLE login (
 id serial PRIMARY KEY,
 user_id integer NOT NULL,
 hash varchar(100) NOT NULL,
 username text UNIQUE NOT NULL,
 FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade
);

COMMIT;
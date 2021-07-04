BEGIN TRANSACTION;

CREATE TABLE login (
 id serial PRIMARY KEY,
 user_id serial NOT NULL,
 hash varchar(100) NOT NULL,
 username text UNIQUE NOT NULL,
 FOREIGN KEY (user_id) REFERENCES users (id)
);

COMMIT;
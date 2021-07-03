-- create a fresh db
\i '/docker-entrypoint-initdb.d/tables/users.sql'
\i '/docker-entrypoint-initdb.d/tables/blogs.sql'
\i '/docker-entrypoint-initdb.d/tables/questions.sql'
\i '/docker-entrypoint-initdb.d/tables/answers.sql'
\i '/docker-entrypoint-initdb.d/tables/stories.sql'
\i '/docker-entrypoint-initdb.d/tables/messages.sql'

-- For testing purposes only. This file will add dummy data
\i '/docker-entrypoint-initdb.d/seed/seed.sql'
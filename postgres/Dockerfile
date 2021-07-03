FROM postgres:13

ADD /tables/ /docker-entrypoint-initdb.d/tables/
ADD /seed/ /docker-entrypoint-initdb.d/seed/
ADD deploy.sql /docker-entrypoint-initdb.d/
FROM mysql:5.7 as builder

ADD database.sql /docker-entrypoint-initdb.d
ADD database_insert.sql /docker-entrypoint-initdb.d

EXPOSE 3306
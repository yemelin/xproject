#! /bin/bash

APP_DB_USER=xproject
APP_DB_PWD=xproject
docker run --rm -d -p 5432:5432 -v xproject-pgdata:/var/lib/postgresql/data --health-cmd pg_isready --health-interval 0.5s -e POSTGRES_USER=${APP_DB_USER} --name xproject-postgres postgres:10-alpine
until [ `docker inspect --format "{{json .State.Health.Status }}" xproject-postgres` = '"healthy"' ];
do 
    echo "waiting for postgres container..."
    sleep 0.5
done
echo "pg ready"
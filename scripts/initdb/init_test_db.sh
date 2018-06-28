#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        login text NOT NULL UNIQUE,
        password text NOT NULL,
        email text NOT NULL
    );
EOSQL
# DB and Migration tool

## DB
* PostgreSQL


## Migration tool
* https://github.com/golang-migrate/migrate

## Migration details:
* https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md
* migrations/ - project root folder for store all migration files
* all migration files is plain sql file
* file name format:
    1. {000xxx}_some_readable_name_describing_content.{up/down}.sql
    1. number prefix and up/down suffix will generate automatic
* empty up/down files generate command example:
```
migrate create -ext sql -seq -dir migrations create_some_table
```
* apply migrate command example:
```
migrate -database "postgres://${APP_DB_USER}:${APP_DB_PWD}@localhost:5432/tablename" -path migrations
```
* migration current status show:
```
migrate -database "postgres://${APP_DB_USER}:${APP_DB_PWD}@localhost:5432/tablename" -path migrations version
```
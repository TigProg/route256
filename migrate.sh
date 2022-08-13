#!/bin/sh

export MIGRATIONS_DIR="migrations"
export DB_DSN="host=localhost port=5432 user=user password=password dbname=bus_booking sslmode=disable"

if [ "$1" = "--dryrun" ]; then
    goose -v -dir ${MIGRATIONS_DIR} postgres "${DB_DSN}" status
elif [ "$1" = "--down" ]; then
    # roll back a single migration from the current version
    goose -v -dir ${MIGRATIONS_DIR} postgres "${DB_DSN}" down
elif [ "$1" = "--test" ]; then
    # apply one transaction, and (if exist) roll back
    goose -v -dir ${MIGRATIONS_DIR} postgres "${DB_DSN}" up-by-one &&
    goose -v -dir ${MIGRATIONS_DIR} postgres "${DB_DSN}" down  &&
    echo "migration up and migration down is done successfully"
else
    # apply all available migrations
    goose -v -dir ${MIGRATIONS_DIR} postgres "${DB_DSN}" up
fi

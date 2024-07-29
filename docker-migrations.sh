#!/bin/bash

DB_DSN="host=$DB_HOST port=$DB_PORT user=$DB_USERNAME password=$DB_PASSWORD dbname=$DB_NAME sslmode=$DB_SSL"

goose -dir ./migrations postgres "$DB_DSN" up
#!/bin/sh

set -e

host="$POSTGRES_HOST"
port="$POSTGRES_PORT"

echo "Waiting for PostgreSQL on $host:$port..."

while ! nc -z "postgres" "5432"; do
  sleep 10
done

echo "PostgreSQL is up - executing command"
exec "$@"
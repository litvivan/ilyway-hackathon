#!/bin/sh

#workdir="../${PWD}"
#go build -o ./bin/build ./cmd/main.go
#./bin/build

cd /app

echo "Waiting db to launch on ${DB_PORT}..."

while ! netcat -z $DB_HOST $DB_PORT; do  
  echo "Waiting for db on port ${DB_PORT}..."
  sleep 1 # wait for 1 second before check again
done

make migrate && make build && make run
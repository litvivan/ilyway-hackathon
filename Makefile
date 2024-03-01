migrate:
	sql-migrate up -config=conf/dbconfig.yml
.PHONY: clean

rollback:
	sql-migrate down -config=conf/dbconfig.yml
.PHONY: migrate

migrate-status:
	sql-migrate status -config=conf/dbconfig.yml
.PHONY: clean

migration-new:
	sql-migrate new -config=conf/dbconfig.yml
.PHONY: clean

gen: gen-sqlc
	go generate ./...
.PHONY: generate

gen-sqlc:
	sqlc generate --file infra/db/sqlc.yml
.PHONY: gen-sqlc

dev-deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/vektra/mockery/v2@v2.42.0
.PHONY: dev-deps

build:
	go build -o ./bin/build ./cmd/main.go
.PHONY: build

run:
	/bin/sh -c "./bin/build"
.PHONY: run

test:
	go test -v --race -p 4 ./...
.PHONY: test
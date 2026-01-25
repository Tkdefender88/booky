.PHONY: all

all: build

build:
	go build -o ./bin/booky main.go

clean:
	rm -rf ./bin

migrate:
	@goose -dir ./internal/repo/migrations/ sqlite3 booky.db up

generate:
	@sqlc generate

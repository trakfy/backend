.PHONY: run format test migrate-up migrate-down migrate-create

# Carregar as variáveis do arquivo .env
include .env
export

run:
	docker-compose -f .trakfy/docker-compose.yml up --build

format:
	go fmt ./...

test:
	go test -v ./...

migrate-up:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose down 1

migrate-create:
	migrate create -ext sql -dir ./db/migrations -seq $(name)
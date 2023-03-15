POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DATABASE=botdb

-include .env

DB_URL="postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable"

print:
	echo "$(DB_URL)"

start:
	go run cmd/main.go
create:
	migrate create -ext sql -dir migrations -seq create_search_trigger_table

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down
fix:
	migrate -path migrations -database "$(DB_URL)" force 2 1

.PHONY: start migrateup migratedown
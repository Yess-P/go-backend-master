postgres:
	docker run -p 5432:5432 --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=7912 -d postgres:16.0

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank 

dropdb:
	docker exec -it postgres dropdb simple_bank 

migrationup:
	migrate -path db/migration -database "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrationup migrationdown sqlc test server
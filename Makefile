postgres:
	docker run -p 5432:5432 --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=7912 -d postgres:16.0

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank 

dropdb:
	docker exec -it postgres dropdb simple_bank 

migrationup:
	migrate -path db/migration -database "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationup1:
	migrate -path db/migration -database "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migrationdown:
	migrate -path db/migration -database "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrationdown1:
	migrate -path db/migration -database "postgresql://root:7912@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	 mockgen -package mockdb  -destination db/mock/store.go  backend-master/db/sqlc Store

add_users:
	migrate create -ext sql -dir db/migration -seq add_users
new_migration:
	migrate create -ext sql -dir db/migration -seq add_session

.PHONY: postgres createdb dropdb migrationup migrationdown migrationup1 migrationdown1 sqlc test server mock add_users new_migration
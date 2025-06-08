postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret --label project=go-dutch -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root go-dutch

dropdb:
	docker exec -it postgres17 dropdb go-dutch

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go-dutch?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go-dutch?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go-dutch?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go-dutch?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server
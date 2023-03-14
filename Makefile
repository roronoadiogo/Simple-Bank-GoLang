postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank_account 
dropdb:
	docker exec -it postgres12 dropdb --username=root bank_account
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/bank_account?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/bank_account?sslmode=disable" -verbose down

test:
	go test -v -cover
	
sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc


	
postgresinit:
	docker run --name markitdb -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

startdb:
	docker start markitdb 

postgres:
	docker exec -it markitdb psql

createdb:
	docker exec -it markitdb createdb --username=root --owner=root markit

dropdb:
	docker exec -it markitdb dropdb markit

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5434/markit?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5434/markit?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown startdb

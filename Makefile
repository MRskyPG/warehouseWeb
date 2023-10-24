run:
	go run ./cmd/main.go

build:
	go mod download
	docker run --name warehouseWeb -e POSTGRES_PASSWORD=qwerty --rm -d -p 5436:5432 postgres
	make migrate


migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate_drop:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
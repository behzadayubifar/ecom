BINARY_NAME=ecom
.DEFAULT_GOAL=run

build:
	go build -o bin/${BINARY_NAME} cmd/main.go

test:
	go test -v ./...

run: build
	./bin/${BINARY_NAME}

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

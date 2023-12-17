DATABASE_URL:="postgres://postgres:postgres@127.0.0.1:5432/auth?sslmode=disable"

.PHONY: start
start:
	@go run cmd/main.go

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o auth *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: sqlc
sqlc:
	@sqlc generate -f ./postgres/sqlc.yml

.PHONY: migrate-up
migrate-up:
	@migrate -database ${DATABASE_URL} -path postgres/migrations up

.PHONY: migrate-down
migrate-down:
	@migrate -database ${DATABASE_URL} -path postgres/migrations down
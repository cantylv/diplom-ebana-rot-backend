include build/.env

vendor:
	@go mod vendor

tidy:
	@go mod tidy

deps: tidy vendor

test:
	@echo "Running tests..."
	@go test ./internal/... -cover -short -count=1

lint:
	@echo "Running golangci-lint..."
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	@golangci-lint run --fix ./...

migrations-apply:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose up

migrations-down:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose down

generate-sql:
	@echo "Generate sql..."
	@go run -mod=mod github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0 generate

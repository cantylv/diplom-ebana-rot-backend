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
	# @GOOSE_DRIVER=${GOOSE_DRIVER} goose -dir migrations postgres up

migrations-down:
	# @"${DOCKER_BUILD_CLI}" goose -dir migrations postgres "${DOCKER_MASTER_DSN}" down

generate-sql:
	@echo "Generate sql..."
	@go run -mod=mod github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0 generate

include build/.env

vendor:
	@go mod vendor

tidy:
	@go mod tidy

up:
	@docker compose -f build/docker-compose.yaml up -d

deps: tidy vendor

lint:
	@echo "Running golangci-lint..."
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	@golangci-lint run --fix ./...

migrations-apply:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose up

docker-migrations-apply:
	docker exec -it song goose up

migrations-down:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose down

docker-migrations-down:
	@docker exec -it song goose down

# make migrations-create NAME=some_name
migrations-create:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -s --dir ./migrations -timeout 2m create $(NAME) sql

docker-infra: deps up



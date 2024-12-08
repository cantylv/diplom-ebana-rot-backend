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

docker-migrations-apply:
	docker exec -it song goose up

docker-migrations-down:
	@docker exec -it song goose down

# make migrations-create NAME=some_name
migrations-create:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -s --dir ./migrations -timeout 2m create $(NAME) sql

docker-swag:
	@docker exec -it -w /home/effective_mobile_backend song swag init --generalInfo ./cmd/main/main.go

docker-infra: deps up

pipeline: docker-migrations-apply docker-swag
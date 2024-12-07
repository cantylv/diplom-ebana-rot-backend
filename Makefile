include Makefile.common.mk

# make migrations-create NAME=some_name
migrations-create:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -s --dir ./migrations -timeout 2m create $(NAME) sql

go-generate:
	@go generate ./internal/...

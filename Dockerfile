# Builder state
FROM golang:1.23.4 AS builder
RUN apt-get update && apt-get install -y make git curl gcc g++ && apt-get clean
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
ARG MODULE_NAME_BUILDER=effective_mobile_backend
WORKDIR /home/${MODULE_NAME_BUILDER}
COPY . .
# building exe file
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main/main.go
RUN chown root:root main
CMD ["./main"]
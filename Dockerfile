# Builder state
FROM golang:1.21.7 AS builder
RUN apt-get update && apt-get install -y make git curl && apt-get clean

ARG MODULE_NAME_BUILDER=effective_mobile_backend
WORKDIR /home/${MODULE_NAME_BUILDER}

COPY . .

# building exe ile
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main/main.go

# Production state
FROM alpine:3.20.3 as production
WORKDIR /root/
ARG MODULE_NAME_PRODUCTION=effective_mobile_backend

COPY --from=builder /home/${MODULE_NAME_PRODUCTION}/config/config.yaml config/config.yaml
COPY --from=builder /home/${MODULE_NAME_PRODUCTION}/main .

RUN chown root:root main

CMD ["./main"]
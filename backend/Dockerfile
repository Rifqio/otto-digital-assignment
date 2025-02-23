FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY .env .
COPY database/migrations ./database/migrations

RUN echo '#!/bin/sh' > /app/entrypoint.sh && \
    echo 'while ! migrate -database "${DATABASE_URL}" -path database/migrations up; do' >> /app/entrypoint.sh && \
    echo '  echo "Migration failed, retrying in 5 seconds..."' >> /app/entrypoint.sh && \
    echo '  sleep 5' >> /app/entrypoint.sh && \
    echo 'done' >> /app/entrypoint.sh && \
    echo './main' >> /app/entrypoint.sh && \
    chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
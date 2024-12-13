FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ctfplatform ./cmd/main.go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/ctfplatform .

EXPOSE 8080

ENV ADMIN_POSTGRES_URL="postgres://user:password@localhost:5432/dbname"
ENV ADMIN_REDIS_URL="redis://localhost:6379"
ENV ADMIN_APP_PORT="8080"

CMD ["./ctfplatform"]

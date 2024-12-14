FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o admin-app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/admin-app .
COPY ./.env /app/config/config.env
#COPY ../migrations /app/migrations/

EXPOSE 8080

CMD ["./admin-app"]

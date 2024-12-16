FROM golang:latest AS build


WORKDIR /go/src/app
COPY . .
RUN go mod tidy &&\
    go mod download &&\
    CGO_ENABLED=0 GOOS=linux go build -o /app/bin/application cmd/main/main.go && \
    chmod +x /app/bin/application


FROM gcr.io/distroless/static

COPY --from=build /app/bin/application /

EXPOSE 8080
EXPOSE 50051
CMD ["/application"]
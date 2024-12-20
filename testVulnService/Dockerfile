FROM golang:bookworm AS build

WORKDIR /opt
WORKDIR /go/src/app
COPY . .
RUN go mod tidy &&\
    #go mod download &&\
    CGO_ENABLED=0 GOOS=linux go build -o /app/bin/application main.go && \
    chmod +x /app/bin/application


FROM gcr.io/distroless/static

COPY --from=build /app/bin/application /

# REST port
EXPOSE 4040
CMD ["/application"]
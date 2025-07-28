FROM golang:latest AS build


WORKDIR /go/src/app
COPY . .
RUN go mod tidy &&\
    #go mod download &&\
    CGO_ENABLED=0 GOOS=linux go build -o /app/bin/application cmd/main.go && \
    chmod +x /app/bin/application


FROM gcr.io/distroless/static

COPY --from=build /app/bin/application /
COPY --from=build /go/src/app/build/config.prod.yaml /build/config.yaml

# gRPC port
EXPOSE 4000
# REST port
EXPOSE 4010
CMD ["/application"]
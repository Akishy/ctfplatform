FROM golang:bookworm AS build

WORKDIR /opt
RUN wget https://github.com/fullstorydev/grpcurl/releases/download/v1.9.2/grpcurl_1.9.2_linux_x86_64.tar.gz && \
    tar -xvf grpcurl_1.9.2_linux_x86_64.tar.gz && \
    wget https://github.com/jqlang/jq/releases/download/jq-1.7.1/jq-linux-amd64 && \
    chmod +x ./jq-linux-amd64
WORKDIR /go/src/app
COPY . .
RUN go mod tidy &&\
    #go mod download &&\
    CGO_ENABLED=0 GOOS=linux go build -o /app/bin/application main.go && \
    chmod +x /app/bin/application && \
    chmod +x /go/src/app/getcheckeruuid.sh


FROM debian:bookworm

WORKDIR /opt
COPY --from=build /opt/grpcurl ./grpcurl
COPY --from=build /opt/jq-linux-amd64 ./jq
COPY --from=build /go/src/app/checkerSystem.proto .
COPY --from=build /go/src/app/getcheckeruuid.sh .

COPY --from=build /app/bin/application .


# REST port
EXPOSE 4030
CMD ["bash", "-c", "sleep 10000000"]

# Определение переменных
PROTO_SRC = ../proto/deployer.proto
PROTO_DIR = ../proto
OUT_DIR = pkg/api/v1

# Компиляция протобуфа
.PHONY: proto-gen
proto-gen:
	protoc -I $(PROTO_DIR) \
       --go_out=$(OUT_DIR) \
       --go-grpc_out=$(OUT_DIR) \
       --go_opt=paths=source_relative \
       --go-grpc_opt=paths=source_relative \
       $(PROTO_SRC)
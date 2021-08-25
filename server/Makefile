PROTOBUF_NAME := shortener
PROTOBUF_SOURCE_FILE := $(PROTOBUF_NAME).proto
PROTOBUF_DIR := $(PROTOBUF_NAME)/
BIN_DIR := bin/
BIN := $(BIN_DIR)server

.PHONY: default
default: deps $(PROTOBUF_DIR)/*.pb.go server

.PHONY: clean
clean:
	rm -rf $(BIN) $(PROTOBUF_DIR)*.go

.PHONY: deps
deps:
	@which protoc-gen-go > /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc > /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: server
server: bin/server
	bin/server

$(BIN): $(BIN_DIR)
	go build -o $@

$(BIN_DIR):
	mkdir -p $@

$(PROTOBUF_DIR)/*.pb.go: $(PROTOBUF_DIR)
	(cd $(PROTOBUF_DIR) && \
	  protoc \
	  --proto_path=../../ \
	  --go_out=. \
	  --go_opt=M$(PROTOBUF_SOURCE_FILE)=./$(PROTOBUF_NAME) \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=. \
	  ../../$(PROTOBUF_SOURCE_FILE) \
	)

$(PROTOBUF_DIR):
	mkdir -p $@


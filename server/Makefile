PROTOBUF_NAME := shortener
PROTOBUF_SOURCE_FILE := $(PROTOBUF_NAME).proto
PROTOBUF_DIR := $(PROTOBUF_NAME)/

.PHONY: default
default: deps $(PROTOBUF_DIR)/*.pb.go server

.PHONY: clean
clean:
	rm -rf $(PROTOBUF_DIR)*

.PHONY: deps
deps:
	@which protoc-gen-go > /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc > /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

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

.PHONY: server
server:
	go run main.go

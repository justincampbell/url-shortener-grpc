PROTOBUF_NAME := shortener
PROTOBUF_SOURCE_FILE := $(PROTOBUF_NAME).proto

.PHONY: default
default: shortener_pb.rb

.PHONY: clean
clean:
	rm -f *_pb.rb

shortener_pb.rb: ../shortener.proto
	  grpc_tools_ruby_protoc \
	  --proto_path=../ \
	  --grpc_out=. \
	  --ruby_out=. \
	  ../$(PROTOBUF_SOURCE_FILE)

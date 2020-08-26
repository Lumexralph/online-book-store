#!/bin/bash

# compile the client
# protoc -I . --go_out=internal/grpc/domain_service:. internal/proto_files/domain/product.proto

protoc -I . --go_out=plugins=grpc:. internal/proto_files/domain/product.proto

# To be able to use the protobufs as domain for Gorm ORM, run the following commands
# in the terminal in the directory where the stub was generated
protoc-go-inject-tag -input=./product.pb.go
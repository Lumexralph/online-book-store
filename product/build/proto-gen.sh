#!/bin/bash

# compile the client
protoc -I . --go_out=internal/grpc/domain_service:. internal/proto_files/domain/product.proto

# try this out 
# protoc calculator.proto --go_out=plugins=grpc:../server/calculatorpb/

protoc -I . --go_out=plugins=grpc:. internal/proto_files/service/product_service.proto

# protoc-go-inject-tag -input=./product.pb.g

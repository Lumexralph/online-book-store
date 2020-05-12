#!/bin/bash

# compile the client
protoc -I . --go_out=internal/grpc/domain_service:. internal/proto_files/domain/product.proto

# compile the client-service
# protoc -I $PWD --go_out=plugins=grpc:$PWD $PWD/internal/proto-files/service/client-service.proto
# protoc -I.  --go_out=plugins=grpc:. internal/proto-files/service/client-service.proto

protoc -I . --go_out=plugins=grpc:. internal/proto_files/service/product_service.proto

#!/bin/bash

# make generates the JS web product service, but still reports error
# protoc --proto_path=../../protobuf store.proto \
#     --js_out=import_style=commonjs:./services \
#     --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./services

protoc --proto_path=../../protobuf store.proto \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:./src/services

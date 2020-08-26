# make generates the JS web product service, but still reports error
protoc --proto_path=../product internal/proto_files/domain/product.proto \
    --js_out=import_style=commonjs:./services \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./services

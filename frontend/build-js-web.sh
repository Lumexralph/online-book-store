# make generates the JS web product service, but still reports error
protoc --proto_path=../product internal/proto_files/service/product_service.proto \
    --js_out=import_style=commonjs:. \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.

# also build the important Product message
protoc -I ../product --js_out=import_style=commonjs,binary:. internal/proto_files/domain/product.proto

# The trouble of separating domain proto from service seems not to worth it.
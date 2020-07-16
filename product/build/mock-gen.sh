# using GoMock, generate mock for the ProductService Client
mockgen product/internal/grpc/domain ProductServiceClient > cmd/grpc/internal/grpc/domain/mock_domain/mocks.go
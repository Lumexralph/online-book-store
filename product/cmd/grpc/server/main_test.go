// Package main is the implementation of the grpc product service server.
package main

import (
	"context"
	"net"
	"product/internal/grpc/domain"
	"product/internal/grpc/services"
	repo "product/internal/repository"
	repomock "product/internal/repository/mocks"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// TODO: The mock test is failing, I should switch to integrated testing
// using an actual datastore or another mock implementation
var mockStore = &repomock.ProductRepository{}

// duck-typing at work
var r repo.ProductRepository = mockStore

// ref: https://stackoverflow.com/questions/42102496/testing-a-grpc-service/52080545#52080545
func startGRPCServer(t *testing.T) (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	// new product service
	srv := services.NewProductService(mockStore)
	// create grpc server
	s := grpc.NewServer()
	// register the ProductService implementation
	domain.RegisterProductServiceServer(s, srv)

	// starting the server will block main goroutine till an error occurs,
	// so we put it an a goroutine.
	go func() {
		if err := s.Serve(listener); err != nil {
			require.Nilf(t, err, "grpc: could not start test database: %v", err)
		}
	}()

	return s, listener
}

func createBufDialer(l *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return l.Dial()
	}
}

func TestProductGRPCServer_AddProduct(t *testing.T) {
	ctx := context.Background()
	s, lis := startGRPCServer(t)
	defer s.Stop()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(createBufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		require.Nilf(t, err, "grpc: could not dial listener: %v", err)
	}
	defer conn.Close()

	prod := &domain.Product{
		Id:          1,
		Name:        "Sherlock Holmes",
		Description: "It is an investigative book",
		Price:       123.56,
		Slug:        "sherlock-holmes",
		Quantity:    60,
		CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
	}
	// when the CreateProduct mock is called
	mockStore.On("CreateProduct", prod).Return(nil)

	client := domain.NewProductServiceClient(conn)
	product := &domain.AddProductRequest{
		Product: prod,
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	//	try creating a product
	resp, err := client.AddProduct(ctx, product)
	if err != nil {
		require.Nilf(t, err, "AddProduct(): could not create product: %v", err)
	}

	assert.Equalf(t, 1, resp.CreatedProduct.Id, "client.AddProduct(ctx, product) want=%d; got=%d", 1, resp.CreatedProduct.Id)
}

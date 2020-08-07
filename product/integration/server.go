// Package integration_test has the implementation of the gRPC server
package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"product/internal/grpc/domain"
	"product/internal/grpc/services"
	repo "product/internal/repository"
)

// fakeProductStore is temporary storage products
// var fakeProductStore = []*domain.Product{}

// TODO: The mock test is failing, I should switch to integrated testing
// using an actual datastore or another mock implementation
// var mockStore = &repomock.ProductRepository{}

type fakeProductRepository struct {
	store []*domain.Product // temporary storage for products
}

func (fp *fakeProductRepository) CreateProduct(p *domain.Product) (*domain.Product, error) {
	// temporary store product
	fp.store = append(fp.store, p)
	return p, nil
}

// To make sure the FakeProductRepository implements the ProductRepository interface
// and to also allow compilation.
var _ repo.ProductRepository = (*fakeProductRepository)(nil)

func startGRPCServer(t *testing.T) (*grpc.Server, *bufconn.Listener) {
	// 1 << 20 == 1024 * 1024 i.e 1MB
	bufferSize := 1024 * 1024
	lis := bufconn.Listen(bufferSize)

	// duck-typing at work
	fakeRepo := &fakeProductRepository{}
	srv := services.NewProductService(fakeRepo)
	s := grpc.NewServer()
	// register the ProductService implementation
	domain.RegisterProductServiceServer(s, srv)

	// starting the server will block main goroutine till an error occurs,
	// so we put it an a goroutine.
	go func() {
		if err := s.Serve(lis); err != nil {
			require.Nilf(t, err, "grpc: could not start test database: %v", err)
		}
	}()

	return s, lis
}

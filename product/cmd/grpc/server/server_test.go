// Package main is the implementation of the grpc product service server.
package main

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
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

type FakeProductRepository struct {
	store []*domain.Product // temporary storage for products
}

func (fp *FakeProductRepository) CreateProduct(p *domain.Product) (*domain.Product, error) {
	// temporary store product
	fp.store = append(fp.store, p)
	return p, nil
}

// duck-typing at work
var fakeRepo = &FakeProductRepository{}

// just making sure the fake implements the ProductRepository interface
var fakeStore repo.ProductRepository = fakeRepo

// ref: https://stackoverflow.com/questions/42102496/testing-a-grpc-service/52080545#52080545
func startGRPCServer(t *testing.T) (*grpc.Server, *bufconn.Listener) {
	// 1 << 20 == 1024 * 1024 i.e 1MB
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)

	srv := services.NewProductService(fakeStore)
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

	assert.Equalf(t, uint64(1), resp.CreatedProduct.Id, "client.AddProduct(ctx, product) want=%d; got=%d", 1, resp.CreatedProduct.Id)

	if diff := cmp.Diff(resp.CreatedProduct, prod, cmpopts.IgnoreTypes(domain.Timestamp{}, domain.Product{}.XXX_sizecache)); diff != "" {
		t.Errorf("ProductService.AddProduct() mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(resp.CreatedProduct, fakeRepo.store[0], cmpopts.IgnoreTypes(domain.Timestamp{}, domain.Product{}.XXX_sizecache)); diff != "" {
		t.Errorf("ProductRepository.CreateProduct() mismatch (-want +got):\n%s", diff)
	}
}

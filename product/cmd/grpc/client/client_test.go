package main

import (
	"context"
	"net"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"product/internal/grpc/domain"
	"product/internal/grpc/services"
	repo "product/internal/repository"
)

type FakeProductRepository struct {
	store []*domain.Product // temporary storage for products
}

func (fp *FakeProductRepository) CreateProduct(p *domain.Product) (*domain.Product, error) {
	// temporary store product
	fp.store = append(fp.store, p)
	return p, nil
}

var fakeRepo = &FakeProductRepository{}

type productService struct{}

func (s productService) AddProduct(ctx context.Context, in *domain.AddProductRequest) (*domain.AddProductResponse, error) {
	return &domain.AddProductResponse{
		CreatedProduct: fakeRepo.store[0],
	}, nil
}

// duck-typing at work
// just making sure the fake implements the ProductRepository interface
var fakeStore repo.ProductRepository = fakeRepo

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

func TestProductServiceClient_AddProduct(t *testing.T) {
	ctx := context.Background()
	s, lis := startGRPCServer(t)
	defer s.Stop()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(createBufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		require.Nilf(t, err, "grpc: could not dial listener: %v", err)
	}
	defer conn.Close()

	p := &domain.Product{
		Id:          1,
		Name:        "Sherlock Holmes",
		Description: "It is an investigative book",
		Price:       123.56,
		Slug:        "sherlock-holmes",
		Quantity:    60,
		CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
	}

	req := &domain.AddProductRequest{
		Product: p,
	}

	c := domain.NewProductServiceClient(conn)
	resp, err := c.AddProduct(ctx, req)
	if err != nil {
		require.Nilf(t, err, "ProductServiceClient.AddProduct(): got err: %v; want nil err", err)
	}

	if diff := cmp.Diff(resp.CreatedProduct, p, cmpopts.IgnoreTypes(domain.Timestamp{}, domain.Product{}.XXX_sizecache)); diff != "" {
		t.Errorf("ProductServiceClient.AddProduct() mismatch (-want +got):\n%s", diff)
	}
}

// This is a mock implementation, it is left for reference

// func TestAddProduct(t *testing.T) {
// 	resp := &domain.AddProductResponse{
// 		CreatedProduct: product,
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockProdServiceClient := domainmock.NewMockProductServiceClient(ctrl)
// 	mockProdServiceClient.
// 		EXPECT().AddProduct(gomock.Any(), req).
// 		Return(resp, nil)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	resp, err := mockProdServiceClient.AddProduct(ctx, req)
// 	require.Nilf(t, err, "ProductServiceClient.AddProduct(): failed to create product: %v", err)

// 	if diff := cmp.Diff(resp.CreatedProduct, product, cmpopts.IgnoreTypes(domain.Timestamp{}, domain.Product{}.XXX_sizecache)); diff != "" {
// 		t.Errorf("ProductServiceClient.AddProduct() mismatch (-want +got):\n%s", diff)
// 	}
// }

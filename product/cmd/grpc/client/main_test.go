package main

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"product/internal/grpc/domain"
	domainmock "product/internal/grpc/domain/mock_domain"
)

var product = &domain.Product{
	Id:          1,
	Name:        "Sherlock Holmes",
	Description: "It is an investigative book",
	Price:       123.56,
	Slug:        "sherlock-holmes",
	Quantity:    60,
	CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
}

var req = &domain.AddProductRequest{
	Product: product,
}

var resp = &domain.AddProductResponse{
	CreatedProduct: product,
}

func TestAddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProdServiceClient := domainmock.NewMockProductServiceClient(ctrl)
	mockProdServiceClient.
		EXPECT().AddProduct(gomock.Any(), req).
		Return(resp, nil)

	testAddProduct(t, mockProdServiceClient)
}

func testAddProduct(t *testing.T, client domain.ProductServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.AddProduct(ctx, req)
	require.Nilf(t, err, "ProductServiceClient.AddProduct(): failed to create product: %v", err)

	if diff := cmp.Diff(resp.CreatedProduct, product, cmpopts.IgnoreTypes(domain.Timestamp{}, domain.Product{}.XXX_sizecache)); diff != "" {
		t.Errorf("ProductServiceClient.AddProduct() mismatch (-want +got):\n%s", diff)
	}
}

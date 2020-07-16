package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"

	"product/internal/grpc/domain"
	domainmock "product/internal/grpc/domain/mock_domain"
)

// TODO: Finish the client test
func TestAddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &domain.AddProductRequest{
		Product: &domain.Product{
			Id:          1,
			Name:        "Sherlock Holmes",
			Description: "It is an investigative book",
			Price:       123.56,
			Slug:        "sherlock-holmes",
			Quantity:    60,
			CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
		},
	}

	mockProdServiceClient := domainmock.NewMockProductServiceClient(ctrl)
	mockProdServiceClient.
		EXPECT().AddProduct(gomock.Any(), req).
		Return(gomock.Any())
}

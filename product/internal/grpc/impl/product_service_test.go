// Package impl has the ProductService grpc server implementation.
package impl

import (
	"context"
	"product/internal/grpc/domain"
	repo "product/internal/repository"
	repomock "product/internal/repository/mocks"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/go-cmp/cmp"
)

var mockStore = &repomock.ProductRepository{}

// duck-typing at work
var r repo.ProductRepository = mockStore

func TestProductService_AddProduct(t *testing.T) {
	type fields struct {
		store repo.ProductRepository
	}

	type args struct {
		ctx context.Context
		in  *domain.AddProductRequest
	}

	product := &domain.Product{
		Id:          1,
		Name:        "Sherlock Holmes",
		Description: "It is an investigative book",
		Price:       123.56,
		Slug:        "sherlock-holmes",
		CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
	}

	// when the CreateProduct mock is called
	mockStore.On("CreateProduct", product).Return(nil)

	// client request
	request := &domain.AddProductRequest{
		Product: product,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response := &domain.AddProductResponse{CreatedProduct: product}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.AddProductResponse
		wantErr bool
	}{
		{
			name:   "create product success",
			fields: fields{store: r},
			args: args{
				ctx: ctx,
				in:  request,
			},
			want: response,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProductService(tt.fields.store)
			got, err := p.AddProduct(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ProductService.AddProduct() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

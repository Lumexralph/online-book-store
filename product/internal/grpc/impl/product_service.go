// Package impl has the ProductService grpc server implementation.
package impl

import (
	"context"
	"fmt"
	domain "product/internal/grpc/domain_service"
	repo "product/internal/repository"
)

// ProductService is an implementation of ProductServiceServer gRPC service.
type ProductService struct {
	// datastore
	store repo.ProductRepository
}

// NewProductService creates a new product service.
func NewProductService(r repo.ProductRepository) *ProductService {
	return &ProductService{store: r}
}

// AddProduct method will create a new product.
func (p *ProductService) AddProduct(ctx context.Context, in *domain.AddProductRequest) (*domain.AddProductResponse, error) {
	// persist the product
	err := p.store.CreateProduct(in)
	if err != nil {
		fmt.Println(err)
	}

	return &domain.AddProductResponse{
		CreatedProduct: nil,
	}, err

}

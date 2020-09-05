// Package impl has the ProductService grpc server implementation.
package services

import (
	"context"
	"productservice/internal/domain"
	repo "productservice/internal/repository"
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
	product, err := p.store.CreateProduct(in.Product)
	if err != nil {
		return nil, err
	}

	return &domain.AddProductResponse{
		CreatedProduct: product,
	}, err

}

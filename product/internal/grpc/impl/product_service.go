// Package impl has the ProductService grpc server implementation.
package impl

import (
	"context"
	"fmt"
	"product/internal/grpc/domain"
	repo "product/internal/repository"
)

// TODO: add protobuf linter and dependency manager
// ref: https://medium.com/stackpulse/grpc-in-practice-directory-structure-linting-and-more-d4d438ac4f86
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
	_, err := p.store.CreateProduct(in.Product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("product created")
	return &domain.AddProductResponse{
		CreatedProduct: nil,
	}, err

}

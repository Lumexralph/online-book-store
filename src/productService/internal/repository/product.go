// Package repository contains the operations to retrieve or interact with product domain.
// It uses the domain-driven design paradigm such that alternative storage implementations
// may be easily interchanged.
// Reference: https://en.wikipedia.org/wiki/Domain-driven_design
package repository

import (
	"product/internal/grpc/domain"
)

// ProductRepository interface that any database/datastore component implementation have to provide.
type ProductRepository interface {
	CreateProduct(*domain.Product) (*domain.Product, error)
}

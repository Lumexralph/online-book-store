// Package postgres is the implementation of the database session
// in use by the server
package postgres

import (
	"errors"
	"fmt"
	"product/internal/grpc/domain"

	"github.com/jinzhu/gorm"
	//  need it for the side effects of gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ProductStore encapsulates a db connection with the operations
type ProductStore struct {
	DB *gorm.DB
}

// NewProductStore will create a new product store object
func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{DB: db}
}

// CreateProduct will take the data from the stored file
// and persist it to the database.
func (s ProductStore) CreateProduct(product *domain.Product) (*domain.Product, error) {
	db := s.DB.Create(product)
	if db.Error != nil {
		return nil, fmt.Errorf("error creating product: %v", db.Error)
	}

	// check that the created value is domain.Product through type assertion
	p, ok := db.Value.(*domain.Product)
	if !ok {
		return nil, errors.New("product was not created")
	}

	return p, nil
}

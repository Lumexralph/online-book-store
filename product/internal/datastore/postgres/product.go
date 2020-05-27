// Package postgres is the implementation of the database session
// in use by the server
package postgres

import (
	"product/internal/grpc/domain"

	"github.com/jinzhu/gorm"
	//  need it for the side effects of gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ProductStore encapsulates a db connection with the operations
type ProductStore struct {
	DB *gorm.DB
}

// CreateProduct will take the data from the stored file
// and persist it to the database
func (p ProductStore) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return nil, nil
}

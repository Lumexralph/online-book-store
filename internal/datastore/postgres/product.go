//Package datastore is the implementation of the database session
// in use by the server
package postgres

import (
	"database/sql"
	domain "product/internal/grpc/domain_service"

	// register driver needed for postgreSQL
	_ "github.com/lib/pq"
)

// ProductStore encapsulates a db connection with the operations
type ProductStore struct {
	DB *sql.DB
}

// CreateProduct will take the data from the stored file
// and persist it to the database
func (p ProductStore) CreateProduct(product *domain.AddProductRequest) error {
	return nil
}

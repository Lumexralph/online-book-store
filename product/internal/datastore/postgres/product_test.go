// Package postgres is the implementation of the database session
// in use by the server
package postgres

import (
	"product/internal/grpc/domain"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductStore_CreateProduct(t *testing.T) {
	// <setup code>, setup test database
	db, teardown, err := ConfigTestDatabase()
	require.Nil(t, err, "db: could not configure test database")
	// end of setup

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		product *domain.Product
	}

	field := fields{DB: db}
	arg := args{
		product: &domain.Product{
			Id:          1,
			Name:        "Sherlock Holmes",
			Description: "It is an investigative book",
			Price:       123.56,
			Quantity:    20,
			Slug:        "sherlock-holmes",
			CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
		},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.Product
	}{
		{
			name:   "successfully create product",
			fields: field,
			args:   arg,
			want: &domain.Product{
				Id:          1,
				Name:        "Sherlock Holmes",
				Description: "It is an investigative book",
				Price:       123.56,
				Quantity:    20,
				Slug:        "sherlock-holmes",
				CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
				UpdatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProductStore(tt.fields.DB)
			got, err := s.CreateProduct(tt.args.product)

			assert.Nilf(t, err, "ProductStore.CreateProduct() error = %v", err)

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreTypes(domain.Timestamp{})); diff != "" {
				t.Errorf("ProductStore.CreateProduct() mismatch (-want +got):\n%s", diff)
			}
		})
	}

	// <tear-down code>
	err = db.DropTable(&domain.Product{}).Error
	require.Nilf(t, err, "teardown: could not drop table: %v", err)
	err = teardown()
	require.Nilf(t, err, "teardown: error closing the database: %v", err)
}

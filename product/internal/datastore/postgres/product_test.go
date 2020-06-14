// Package postgres is the implementation of the database session
// in use by the server
package postgres

import (
	"product/internal/grpc/domain"
	"testing"

	"github.com/golang/protobuf/ptypes"
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
			t.Log("got", got)
			assert.Nilf(t, err, "ProductStore.CreateProduct() error = %v", err)

			assert.Equalf(t, got.Id, tt.want.Id, "ProductStore.CreateProduct() got=%d; want=%d", got.Id, tt.want.Id)
			assert.Equalf(t, got.Name, tt.want.Name, "ProductStore.CreateProduct() got=%d; want=%d", got.Name, tt.want.Name)
			assert.Equalf(t, got.Description, tt.want.Description, "ProductStore.CreateProduct() got=%d; want=%d", got.Description, tt.want.Description)
			assert.Equalf(t, got.Price, tt.want.Price, "ProductStore.CreateProduct() got=%d; want=%d", got.Price, tt.want.Price)
			assert.Equalf(t, got.Slug, tt.want.Slug, "ProductStore.CreateProduct() got=%d; want=%d", got.Slug, tt.want.Slug)
			assert.Equalf(t, got.CreatedAt, tt.want.CreatedAt, "ProductStore.CreateProduct() got=%d; want=%d", got.CreatedAt, tt.want.CreatedAt)
		})
	}

	// <tear-down code>
	err = db.DropTable(&domain.Product{}).Error
	require.Nilf(t, err, "teardown: could not drop table: %v", err)
	err = teardown()
	require.Nilf(t, err, "teardown: error closing the database: %v", err)

}

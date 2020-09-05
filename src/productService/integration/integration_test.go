package integration

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"productservice/internal/domain"
)

func TestProductService(t *testing.T) {
	s, lis := startGRPCServer(t)
	defer s.Stop()

	conn := startGRPCClient(t, lis)
	defer conn.Close()

	c := domain.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	p := &domain.Product{
		Id:          1,
		Name:        "Sherlock Holmes",
		Description: "It is an investigative book",
		Price:       123.56,
		Slug:        "sherlock-holmes",
		Quantity:    60,
		CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
	}
	req := &domain.AddProductRequest{
		Product: p,
	}

	resp, err := c.AddProduct(ctx, req)
	if err != nil {
		require.Nilf(t, err, "ProductService.AddProduct(): got err: %v; want nil err", err)
	}

	if diff := cmp.Diff(resp.CreatedProduct, p, cmpopts.IgnoreTypes(domain.Timestamp{}, domain.Product{}.XXX_sizecache)); diff != "" {
		t.Errorf("ProductService.AddProduct() mismatch (-want +got):\n%s", diff)
	}
}

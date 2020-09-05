package integration

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func createBufDialer(l *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return l.Dial()
	}
}

func startGRPCClient(t *testing.T, l *bufconn.Listener) *grpc.ClientConn {
	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(createBufDialer(l)), grpc.WithInsecure())
	if err != nil {
		require.Nilf(t, err, "grpc: could not dial listener: %v", err)
	}
	return conn
}

// This is a mock implementation, it is left for reference

// func TestAddProduct(t *testing.T) {
// 	resp := &domain.AddProductResponse{
// 		CreatedProduct: product,
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockProdServiceClient := domainmock.NewMockProductServiceClient(ctrl)
// 	mockProdServiceClient.
// 		EXPECT().AddProduct(gomock.Any(), req).
// 		Return(resp, nil)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	resp, err := mockProdServiceClient.AddProduct(ctx, req)
// 	require.Nilf(t, err, "ProductServiceClient.AddProduct(): failed to create product: %v", err)

// 	if diff := cmp.Diff(resp.CreatedProduct, product, cmpopts.IgnoreTypes(domain.Timestamp{}, domain.Product{}.XXX_sizecache)); diff != "" {
// 		t.Errorf("ProductServiceClient.AddProduct() mismatch (-want +got):\n%s", diff)
// 	}
// }

// Package main contains the implementation of the grpc client
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"productservice/internal/domain"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
)

const address = ":5001"

func uploadProductImage(ctx context.Context, conn *grpc.ClientConn, productID uint64, imagePath string) (*domain.UploadImageResponse, error) {
	// open the image file
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}
	defer f.Close()

	productImageClient := domain.NewProductImageServiceClient(conn)
	stream, err := productImageClient.UploadImage(ctx)
	if err != nil {
		return nil, fmt.Errorf("upload-image: failed to upload image: %v", err)
	}
	// create the first request to send some image information to the server
	req := &domain.UploadImageRequest{
		Data: &domain.UploadImageRequest_Info{
			Info: &domain.ImageInfo{
				ProductId: productID,
				ImageFormat: filepath.Ext(imagePath),
			},
		},
	}

	// send the first request - image metadata
	// NOTE: when an error occurs, the server will close the stream,
	// and thus the client cannot send more data to it.
	// stream.RecvMsg(nil) - gives us more information about the gRPC error
	err = stream.Send(req)
	if err != nil {
		return nil, fmt.Errorf("upload-image: failed to send request to server: %v - %v", err, stream.RecvMsg(nil))
	}

	// create a reader from the image file and send in chunks using the buffer
	r := bufio.NewReader(f)
	buf := make([]byte, 1 << 10)

	// start sending the image data in chunks / partitions
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			// done sending all the image data from file
			break
		}
		if err != nil {
			return nil, fmt.Errorf("upload-image: failed to read image chunks to buffer: %v", err)
		}

		req := &domain.UploadImageRequest{
			Data: &domain.UploadImageRequest_ImageChunk{
				ImageChunk: buf[:n], // keep reading the newly added data
			},
		}

		err = stream.Send(req)
		if err != nil {
			return nil, fmt.Errorf("upload-image: failed to send image chunk to server: %v - %v", err, stream.RecvMsg(nil))
		}
	}

	// get the response from the server
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("upload-image: failed to get response from server: %v - %v", err, stream.RecvMsg(nil))
	}
	return resp, nil
}

func main() {
	// make a grpc server connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	productClient := domain.NewProductServiceClient(conn)
	req := &domain.AddProductRequest{
		Product: &domain.Product{
			Id:          1,
			Name:        "Sherlock Holmes",
			Description: "It is an investigative book",
			Price:       123.56,
			Slug:        "sherlock-holmes",
			Quantity:    60,
			CreatedAt:   &domain.Timestamp{Timestamp: ptypes.TimestampNow()},
		},
	}
	// wait for 5 secs to handle the operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//	try creating a product
	prodresp, err := productClient.AddProduct(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("product-service: successfully created product: ", prodresp)

	imgresp, err := uploadProductImage(ctx, conn, prodresp.CreatedProduct.Id, "tmp/multi-tenant.png")

	log.Println("image-service: successfully uploaded image: ", imgresp)
}

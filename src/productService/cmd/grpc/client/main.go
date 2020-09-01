// Package main contains the implementation of the grpc client
package main

import (
	"context"
	"log"
	"time"

	"product/internal/grpc/domain"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
)

const address = ":5001"

func main() {
	// make a grpc server connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := domain.NewProductServiceClient(conn)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//	try creating a product
	resp, err := client.AddProduct(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("response: ", resp)
}

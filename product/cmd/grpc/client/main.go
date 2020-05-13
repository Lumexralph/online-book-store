// Package main contains the implementation of the grpc client
package main

import (
	"context"
	"log"
	"time"

	domain "product/internal/grpc/domain_service"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
)

const address = ":5000"

func main() {
	// make a grpc server connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := domain.NewProductServiceClient(conn)

	product := &domain.AddProductRequest{
		Product: &domain.Product{
			Id:          "1",
			Name:        "Sherlock Holmes",
			Description: "It is an investigative book",
			Price:       123.56,
			Slug:        "sherlock-holmes",
			Active:      true,
			InStock:     true,
			DateUpdated: ptypes.TimestampNow(),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//	try creating a product
	resp, err := client.AddProduct(ctx, product)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("response: ", resp)
}

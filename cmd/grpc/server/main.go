// Package main is the implementation of the grpc product service server.
package main

import (
	"net"
	"fmt"
	"os"
	"log"
	"google.golang.org/grpc"
	"product/internal/datastore/postgres"
	service "product/internal/grpc/impl"
	domain "product/internal/grpc/domain_service"

	_ "github.com/joho/godotenv/autoload"
)

const port = ":5000"

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	host := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// create database url
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		dbport,
		user,
		password,
		dbname,
		sslmode,
	)
	db, teardown, err := postgres.Connect(connStr)
	if err != nil {
		return err
	}
	defer teardown()

	// product service store
	prodStore := postgres.ProductStore{
		DB: db,
	}
	// new product service
	srv := service.NewProductService(prodStore)
	// create grpc server
	s := grpc.NewServer()
	// register the ProductService implementation
	domain.RegisterProductServiceServer(s, srv)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	log.Printf("product-service: listening on port %s\n", port)
	return s.Serve(lis)
}

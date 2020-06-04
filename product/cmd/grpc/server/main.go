// Package main is the implementation of the grpc product service server.
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"product/internal/datastore/migrations"
	"product/internal/datastore/postgres"
	"product/internal/grpc/domain"
	service "product/internal/grpc/impl"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const port = ":5001"

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

// run - where the world begins!
func run() error {
	godotenv.Load("../../../.env")
	host := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	// create database url
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		dbport,
		dbuser,
		passwd,
		dbname,
		sslmode,
	)
	db, teardown, err := postgres.Connect(connStr)
	if err != nil {
		return err
	}
	log.Println("postres: database connected!")
	defer teardown()

	// run migrations
	err = migrations.Migrate(db)
	if err != nil {
		log.Fatalf("failed to run migration: %v\n", err)
	}

	// product service store
	prodStore := postgres.NewProductStore(db)
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

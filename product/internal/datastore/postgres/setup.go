// Package testdbsetup is an implementation for setting up
// a test database for the integration tests.
package postgres

import (
	"fmt"
	"log"
	"os"
	"product/internal/datastore/migrations"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// ConfigTestDatabase will handle the configuration of a test database.
func ConfigTestDatabase() (db *gorm.DB, dbTeardown func() error, err error) {
	godotenv.Load("../../../.env")
	host := os.Getenv("TESTDB_HOST")
	dbport := os.Getenv("TESTDB_PORT")
	dbuser := os.Getenv("TESTDB_USER")
	passwd := os.Getenv("TESTDB_PASSWORD")
	dbname := os.Getenv("TESTDB_NAME")
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
	db, dbTeardown, err = Connect(connStr)
	if err != nil {
		return nil, nil, err
	}
	log.Println("postres: test database connected!")

	// run migrations
	err = migrations.Migrate(db)
	if err != nil {
		log.Fatalf("failed to run migration: %v\n", err)
	}
	return
}

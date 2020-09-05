// Package migrations has the entities that created as
// database models from the genrated protobuf which is
// persisted as tables in the database.
package migrations

import (
	"log"

	"productservice/internal/domain"

	"github.com/jinzhu/gorm"
	//  need it for the side effects of gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Migrate will create a new database tables with the generated protobuf entities.
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.Product{}, &domain.Category{}).Error
	if err != nil {
		return err
	}
	log.Println("postgres: migration done!")
	return nil
}

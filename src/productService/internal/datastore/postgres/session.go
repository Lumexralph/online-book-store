package postgres

import (
	"github.com/jinzhu/gorm"
	//  need it for the side effects of gorm for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect will create a new database connection with the supplied psqlInfo
func Connect(psqlInfo string) (*gorm.DB, func() error, error) {
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, nil, err
	}
	return db, db.Close, nil
}

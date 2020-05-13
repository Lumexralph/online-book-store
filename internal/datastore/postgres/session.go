package postgres

import "database/sql"

// DBConnection will create a new database connection with the supplied psqlInfo
func Connect(psqlInfo string) (*sql.DB, func() error, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, nil, err
	}
	return db, db.Close, nil
}

package database

import (
	"api/internal/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect open database connection
func Connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}

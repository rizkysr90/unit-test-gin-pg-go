package db

import (
	"database/sql"
)

func InitDB() (*sql.DB, error) {
	// You should store your credentials in environment variable, not hardcoded like this
	connStr := "postgres://postgres:mysecretpassword@localhost/example?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

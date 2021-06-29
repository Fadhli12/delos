package config

import (
	"database/sql"
	"fmt"
)

var PostgreConn *sql.DB

func PostgreConnection(config *PostgreSQL) (*sql.DB, error) {
	connection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.USER,
		config.PASSWORD,
		config.HOST,
		config.PORT,
		config.DATABASE_NAME)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	PostgreConn = db
	return db, nil
}

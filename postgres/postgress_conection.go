package postgres

import (
	"database/sql"
	"fmt"
	"imtixon5/config"

	_ "github.com/lib/pq"
)

func ConnectionDb() (*sql.DB, error) {
	conf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_NAME, conf.DB_PASSWORD)
	db, err := sql.Open("postgres", conDb)
	if err != nil {
			return nil, err
	}

	if err := db.Ping(); err != nil {
			return nil, err
	}

	return db, nil
}
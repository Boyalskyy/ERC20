package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	host     = "db"
	port     = "5432"
	user     = "postgres"
	password = "qwe1asd1"
	dbname   = "postgres"
	sslmode  = "disable"
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode))
	if err != nil {
		return nil, errors.Wrap(err, "DB connection error")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Ping DB connection error")
	}
	return db, nil
}

package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "qwe1asd1"
	dbname   = "postgres"
	sslmode  = "disable"
)

//var schema = `
//CREATE TABLE events (
//    id SERIAL PRIMARY KEY,
//    LogName TEXT,
//    From TEXT,
//    To TEXT,
//    Amount TEXT
//);`

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode))
	if err != nil {
		return nil, errors.Wrap(err, "DB connection error")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "DB connection error")
	}
	//db.MustExec(schema)
	return db, nil
}

func RecordingEvents(db *sqlx.DB, from string, to string, amount string, err error) {

}

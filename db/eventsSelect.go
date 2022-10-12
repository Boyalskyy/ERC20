package db

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const tableName = "events"

var eventsSelect = sq.Select("r.*").From(fmt.Sprintf("%s r", tableName)).PlaceholderFormat(sq.Dollar)

type EventsQuery struct {
	db  *sqlx.DB
	sql sq.SelectBuilder
}

func NewEvents(db *sqlx.DB) *EventsQuery {
	return &EventsQuery{
		db:  db,
		sql: eventsSelect,
	}
}

type Events struct {
	ID          int64  `db:"id"`
	LogName     string `db:"log_name"`
	AddressFrom string `db:"address_from"`
	AddressTo   string `db:"address_to"`
	Amount      string `db:"amount"`
}
